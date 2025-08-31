package service

//
//import (
//	"github.com/go-kratos/kratos/v2/transport/http"
//)
//
//type CreateNewsReq struct {
//	Title   string `json:"title"`
//	Content string `json:"content"`
//}
//
//type CreateNewsResp struct {
//	Id int32 `json:"id"`
//}
//
//func CreateNews(ctx http.Context) error {
//	//Todo
//
//	ctx.JSON(200, CreateNewsResp{
//		Id: 1,
//	})
//	return nil
//}
//
//type GetNewsDetailReq struct {
//	Id int32 `json:"id"`
//}
//
//type GetNewsDetailResp struct {
//	Id         int32  `json:"id"`
//	Title      string `json:"title"`
//	Content    string `json:"content"`
//	CreateTime string `json:"create_time"`
//}
//
//func GetNewsDetailNews(ctx http.Context) error {
//	//TODO
//
//	ctx.JSON(200, GetNewsDetailResp{
//		Id: 1,
//	})
//	return nil
//}
