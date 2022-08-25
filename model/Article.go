package model

import (
	"goblog/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category   Category `gorm:"foreignkey:CategoryId"`
	Title      string   `gorm:"type varchar(100);not null;comment:文章标题" json:"title"`
	CategoryId int      `gorm:"type int;not null;comment:分类id" json:"category_id"` // CategoryId 必须写成这样，否则这里的外键对应不上Category的主键id
	Desc       string   `gorm:"type varchar(200);comment:文章描述" json:"desc"`
	Content    string   `gorm:"type longtext;comment:文章内容" json:"content"`
	Img        string   `gorm:"type varchar(100);comment:文章图片" json:"img"`
}

// 新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCSE
}

// 查询分类下所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64

	err := db.Preload("Category").Where("category_id = ?", id).Find(&cateArtList).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCSE, total
}

// 查询文章列表
func GetArt(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articlelist []Article
	var total int64
	// 分页显示
	if title == "" {
		err = db.Order("Updated_At DESC").Preload("Category").Find(&articlelist).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
		// 单独计数
		db.Model(&articlelist).Count(&total)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errmsg.ERROR, 0
		}
		return articlelist, errmsg.SUCCSE, total
	}
	err = db.Order("Updated_At DESC").Preload("Category").Where("title LIKE ?", title+"%").Find(&articlelist).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	// 单独计数
	db.Model(&articlelist).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articlelist, errmsg.SUCCSE, total
}

// 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["category_id"] = data.CategoryId
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除文章
func DeleteArt(id int) int {
	var art Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
