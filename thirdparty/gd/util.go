package gd

import (
	"bytes"
	"io"
	"strings"

	"github.com/bogem/id3v2"
	"github.com/go-flac/flacpicture/v2"
	"github.com/go-flac/flacvorbis/v2"
	goflac "github.com/go-flac/go-flac/v2"
	"github.com/pkg/errors"
)

type MusicMetaDataHandler interface {
	AddCommentMetaData() error
	AddLyric(lyric string) error
	AddPic(picReader io.Reader, picExtension string) error
	toReader() (io.Reader, error)
	Close() error
}

type FlacMusicMetaDataHandler struct {
	Music         Music
	flacFile      *goflac.File
	vorbisComment *flacvorbis.MetaDataBlockVorbisComment
}

var _ MusicMetaDataHandler = (*FlacMusicMetaDataHandler)(nil)

func NewFlacMusicMetaDataHandler(music Music, reader io.Reader) (*FlacMusicMetaDataHandler, error) {
	flacFile, err := goflac.ParseBytes(reader)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &FlacMusicMetaDataHandler{
		Music:         music,
		flacFile:      flacFile,
		vorbisComment: flacvorbis.New(),
	}, nil
}

func (h *FlacMusicMetaDataHandler) AddCommentMetaData() error {
	var err error
	h.vorbisComment = flacvorbis.New()
	err = h.vorbisComment.Add(flacvorbis.FIELD_TITLE, h.Music.Name)
	err = h.vorbisComment.Add(flacvorbis.FIELD_ALBUM, h.Music.Album)
	for _, artist := range h.Music.Artist {
		err = h.vorbisComment.Add(flacvorbis.FIELD_ARTIST, artist)
	}
	err = h.vorbisComment.Add("SOURCE", h.Music.Source)
	err = h.vorbisComment.Add("MUSIC_ID", string(h.Music.Id))
	if err != nil {
		err = errors.WithStack(err)
	}
	return err
}

func (h *FlacMusicMetaDataHandler) AddLyric(lyric string) error {
	err := h.vorbisComment.Add("LYRICS", lyric)
	if err != nil {
		err = errors.WithStack(err)
	}
	return err
}

func (h *FlacMusicMetaDataHandler) AddPic(picReader io.Reader, picExtension string) error {
	// 创建图片元数据块
	if picExtension != "jpeg" && picExtension != "png" {
		return errors.Errorf("not valid pic format: %s\n", picExtension)
	}
	pic, err := io.ReadAll(picReader)
	if err != nil {
		return errors.WithStack(err)
	}
	picture, err := flacpicture.NewFromImageData(
		flacpicture.PictureTypeFrontCover,
		"Front cover",
		pic,
		"image/"+picExtension,
	)
	if err != nil {
		return errors.WithStack(err)
	}
	picturemeta := picture.Marshal()
	h.flacFile.Meta = append(h.flacFile.Meta, &picturemeta)
	return nil
}

func (h *FlacMusicMetaDataHandler) toReader() (io.Reader, error) {
	var buffer bytes.Buffer
	cmtsmeta := h.vorbisComment.Marshal()
	h.flacFile.Meta = append(h.flacFile.Meta, &cmtsmeta)
	_, err := h.flacFile.WriteTo(&buffer)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &buffer, nil
}

func (h *FlacMusicMetaDataHandler) Close() error {
	return h.flacFile.Close()
}

type Mp3MusicMetaDataHandler struct {
	Music Music
	tag   *id3v2.Tag
}

var _ MusicMetaDataHandler = (*Mp3MusicMetaDataHandler)(nil)

func NewMp3MusicMetaDataHandler(music Music, reader io.Reader) (*Mp3MusicMetaDataHandler, error) {
	tag, err := id3v2.ParseReader(reader, id3v2.Options{Parse: true})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	tag.SetVersion(4)
	return &Mp3MusicMetaDataHandler{
		Music: music,
		tag:   tag,
	}, nil
}

func (h *Mp3MusicMetaDataHandler) AddCommentMetaData() error {
	h.tag.SetTitle(h.Music.Name)
	artistsFrame := id3v2.TextFrame{
		Encoding: id3v2.EncodingUTF8,
		Text:     strings.Join(h.Music.Artist, "\n"),
	}
	h.tag.AddFrame(h.tag.CommonID("Artists"), artistsFrame)
	h.tag.SetAlbum(h.Music.Album)
	return nil
}

func (h *Mp3MusicMetaDataHandler) AddLyric(lyric string) error {
	lyricsFrame := id3v2.UnsynchronisedLyricsFrame{
		Encoding:          id3v2.EncodingUTF8,
		Language:          "chi",
		ContentDescriptor: "Lyrics",
		Lyrics:            lyric,
	}
	h.tag.AddUnsynchronisedLyricsFrame(lyricsFrame)
	return nil
}

func (h *Mp3MusicMetaDataHandler) AddPic(picReader io.Reader, picExtension string) error {
	pic, err := io.ReadAll(picReader)
	if err != nil {
		return errors.WithStack(err)
	}
	cover := id3v2.PictureFrame{
		Encoding:    id3v2.EncodingUTF8,
		MimeType:    "image/" + picExtension,
		PictureType: id3v2.PTFrontCover,
		Description: "Front cover",
		Picture:     pic,
	}
	h.tag.AddAttachedPicture(cover)
	return nil
}

func (h *Mp3MusicMetaDataHandler) toReader() (io.Reader, error) {
	var buffer bytes.Buffer
	_, err := h.tag.WriteTo(&buffer)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &buffer, nil
}

func (h *Mp3MusicMetaDataHandler) Close() error {
	return h.tag.Close()
}
