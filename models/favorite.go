package models

import(
	"im/config"
	"strconv"
	"fmt"
	_ "github.com/gin-gonic/gin"
)

type Favorite struct {

	Identity   uint       `json:"identity"`

	VideoId    uint		  `json:"video_id"`

	Status	   uint

}


func FavoriteList(userId uint) (PublishListResponse, error) {
	//查询出视频id信息
	videoBaseList, err := SelectVideoInfoListByUserId(int64(userId))
	if err != nil {
		return PublishListResponse{
			StatusCode: 1,
			StatusMsg:  "查询点赞视频id操作失败",
			UserList:   nil,
		}, err
	}

	videos := make([]Video, 0)

	for _, video := range videoBaseList{
		author, _ := GetUserBasicByAccount(strconv.Itoa(int(userId)))
		videos = append(videos, Video{
			Author:        *author,
			CommentCount:  config.CommentCount,
			CoverUrl:      config.CoverUrl,
			FavoriteCount: config.FavoriteCount,
			Identity:      video.Identity,
			PlayUrl:       config.PlayUrl,
			Title:         config.Tiele,
		})
	}
	return PublishListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		UserList:   videos,
	}, nil
}

// 点赞操作
func FavoriteAction(userId uint, videoId uint, actionType uint) (err error) {
	//点赞
	if actionType == 1 {
		var favoriteExit = &Favorite{} //没有时会返回错误
		sqlStr := "select user_id,video_id from favorites user_id = ? AND video_id = ? "
		err := db.Select(&favoriteExit, sqlStr, userId, videoId)
		if err != nil {
			favoritesql := "insert into favorites(user_id, video_id) values(?,?)" //插入favorite数据
			ret, err := db.Exec(favoritesql, userId, videoId)
			if err != nil {
				fmt.Printf("insert error, err:%v\n", err)
				return nil
			}
			theID, err := ret.LastInsertId() //新插入数据的id
			if err != nil {
				fmt.Printf("inert ID failed, err :%v\n", err)
				return nil
			}
			fmt.Printf("insert success,the id is %d. \n", theID)
			favoritevideo := "update video_infos set favorite_count =favorite_count+1 where author_id = ?"
			reta, err := db.Exec(favoritevideo, videoId)
			if err != nil {
				fmt.Printf("update failed,err:%v \n", err)
				return nil
			}
			ne, err := reta.RowsAffected() //操作影响的行数
			if err != nil {
				fmt.Printf("get RowAffected failed, err: %v \n", err)
				return nil
			}
			fmt.Printf("update success ,affect rows:%d \n", ne)
			//userId的favorite_count增加
			if err := AddFavorite(userId); err != nil {
				return err
			}
			//videoId对应的userId的喜欢数量增加
			GuId, err := GetVideoAuthor(videoId)
			if err != nil {
				return err
			}
			if err := AddTotalFavorited(GuId); err != nil {
				return err
			}
		} else { //如果存在
			if favoriteExit.Status == 0 { //status为0-video的favorite_count加1
				sqlStr := "update video_infos set favorite_count=favorite_count+1 where author_id =?"
				ret, err := db.Exec(sqlStr, videoId)
				if err != nil {
					fmt.Printf("update failed, err:%v \n", err)
					return nil
				}
				ne, err := ret.RowsAffected()
				if err != nil {
					fmt.Printf("get RowsAffected failed, err: %v\n", err)
					return nil
				}
				fmt.Printf("update success ,affected rows:%d \n", ne)
				sqlstr := "update favorites set status = 1 where video_id = ?"
				re, err := db.Exec(sqlstr, videoId)
				if err != nil {
					fmt.Printf("update failed, err:%v \n", err)
					return nil
				}
				ns, err := re.RowsAffected()
				if err != nil {
					fmt.Printf("get RowsAffected failed, err:%v \n", err)
					return nil
				}
				fmt.Printf("update success,affected rows:%d\n", ns)
				if err := AddFavorite(userId); err != nil {
					return err
				}
				GueId, err := GetVideoAuthor(videoId)
				if err != nil {
					return err
				}
				if err := AddTotalFavorited(GueId); err != nil {
					return err
				}
			}
			//status为1-video的favorite_count不变
			return nil
		}
	} else { //取消点赞
		var favoriteCancel = &Favorite{}
		sqlStr := "select user_id,video_id from favorites user_id = ? AND video_id = ? "
		err := db.Select(&favoriteCancel, sqlStr, userId, videoId)
		if err != nil { //找不到这条记录，取消点赞失败，创建记录
			favoritesql := "insert into favorites(user_id, video_id) values(?,?)" //插入favorite数据
			ret, err := db.Exec(favoritesql, userId, videoId)
			if err != nil {
				fmt.Printf("insert error, err:%v\n", err)
				return nil
			}
			theID, err := ret.LastInsertId() //新插入数据的id
			if err != nil {
				fmt.Printf("inert ID failed, err :%v\n", err)
				return nil
			}
			fmt.Printf("insert success,the id is %d. \n", theID)
			if err := ReduceFavoriteCount(userId); err != nil {
				return err
			}
			GusId, err := GetVideoAuthor(videoId)
			if err != nil {
				return err
			}
			if err := ReduceTotalFavorited(GusId); err != nil {
				return err
			}
			return err
		}
		//存在
		if favoriteCancel.Status == 1 { //status为1-video的favorite_count减1
			sqlStr := "update video_infos set favorite_count=favorite_count-1 where author_id =?"
			ret, err := db.Exec(sqlStr, videoId)
			if err != nil {
				fmt.Printf("update failed, err:%v \n", err)
				return nil
			}
			ne, err := ret.RowsAffected()
			if err != nil {
				fmt.Printf("get RowsAffected failed, err: %v\n", err)
				return nil
			}
			fmt.Printf("update success ,affected rows:%d \n", ne)
			sqlstr := "update favorites set status = 0 where video_id = ?"
			re, err := db.Exec(sqlstr, videoId)
			if err != nil {
				fmt.Printf("update failed, err:%v \n", err)
				return nil
			}
			ns, err := re.RowsAffected()
			if err != nil {
				fmt.Printf("get RowsAffected failed, err:%v \n", err)
				return nil
			}
			fmt.Printf("update success,affected rows:%d\n", ns)
			if err := ReduceFavoriteCount(userId); err != nil {
				return err
			}
			Ges, err := GetVideoAuthor(videoId)
			if err := ReduceTotalFavorited(Ges); err != nil {
				return err
			}
			return err
		}
		return nil
	}
	return nil
}

// 添加喜欢数量
func AddFavorite(hostId uint) error {
	sqlStr := "update user_infos set favorite_count =favorite_count+1 where id = ?"
	ret, err := db.Exec(sqlStr, hostId)
	if err != nil {
		fmt.Printf("update failed,err:%v \n", err)
		return nil
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed,err:%v \n", err)
		return nil
	}
	fmt.Printf("update succes,affected rows:%d \n", n)
	return nil
}

// 得到视频作者id
func GetVideoAuthor(video uint) (uint, error) {
	sqlStr := "select author_id from video_infos where author_id= ?"
	var u VideoInfo
	err := db.Get(&u, sqlStr, video)
	if err != nil {
		fmt.Printf("get failed,err:%v\n", err)
		return video, nil
	}
	return u.Id, nil
}

// 增加用户的总的喜欢数量
func AddTotalFavorited(hostId uint) error {
	sqlStr := "update user_infos set favorite_count=favorite_count+1 where id = ?"
	ret, err := db.Exec(sqlStr, hostId)
	if err != nil {
		fmt.Printf("update failed,err:%v \n", err)
		return nil
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected faild, err:%v \n", err)
		return nil
	}
	fmt.Printf("update success,affect rows:%d\n", n)
	return nil
}

// 减少喜欢数量
func ReduceFavoriteCount(hostId uint) error {
	sqlStr := "update user_infos set favorite_count=favorite_count-1 where id = ?"
	ret, err := db.Exec(sqlStr, hostId)
	if err != nil {
		fmt.Printf("update failed, err:%v \n", err)
		return nil
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v \n", err)
		return nil
	}
	fmt.Printf("update success, affected rows:%d\n", n)
	return nil
}

// 减少总的喜欢数量
func ReduceTotalFavorited(hostId uint) error {
	sqlStr := "update user_infos set total_favorited=total_favorited-1 where id = ?"
	ret, err := db.Exec(sqlStr, hostId)
	if err != nil {
		fmt.Printf("update failed, err:%v \n", err)
		return nil
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v \n", err)
		return nil
	}
	fmt.Printf("update success, affected rows:%d\n", n)
	return nil
}
