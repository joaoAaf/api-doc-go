package service

import (
	"api-doc-go/dto"
	"api-doc-go/entity"
	"api-doc-go/repository"
	"strconv"
)

func AddId() int {
	repository.LastId += 1
	return repository.LastId
}

func ViewAlbums() []entity.Album {
	return repository.Albums
}

func InsertAlbum(album entity.Album) {
	repository.Albums = append(repository.Albums, album)
}

func UpdateAlbum(album *entity.Album, albumDto dto.AlbumDTO) {
	album.Title = albumDto.Title
	album.Artist = albumDto.Artist
	album.Price = albumDto.Price
}

func DeleteAlbum(index int) {
	repository.Albums = append(repository.Albums[:index], repository.Albums[index+1:]...)
}

func ConvertAlbum(albumDto dto.AlbumDTO) entity.Album {
	var album entity.Album
	album.Id = AddId()
	UpdateAlbum(&album, albumDto)
	return album
}

func VerifyId(strId string) (*entity.Album, int) {
	var album *entity.Album
	index := -1
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(repository.Albums); i++ {
		if repository.Albums[i].Id == id {
			album = &repository.Albums[i]
			index = i
		}
	}
	return album, index
}
