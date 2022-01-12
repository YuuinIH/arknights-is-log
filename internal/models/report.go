package models

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/YuuinIH/is-log/internal/config"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Brief struct {
	Duration int64  `json:"duration"`
	Endat    int64  `json:"endAt"`
	Squad    string `json:"squad"`
	Ending   string `json:"ending"`
}
type Initial struct {
	RecruitGroup string   `json:"recruitGroup"`
	Recruits     []string `json:"recruits"`
	Support      string   `json:"support"`
}
type Zone struct {
	EnterZone  string   `json:"enterZone"`
	Variations []string `json:"variations"`
	NodeList   []Node   `json:"nodeList"`
}
type Node struct {
	Type        int      `json:"type"`
	Stage       string   `json:"stage"`
	Collections []string `json:"collections"`
	Select      []string `json:"select"`
	Capsules    []string `json:"capsules"`
	Tools       []string `json:"tools"`
	Tickets     []string `json:"tickets"`
	Recruits    []string `json:"recruits"`
	Upgrades    []string `json:"upgrades"`
	Shop        Shop     `json:"shop"`
}
type Shop struct {
	Buys   []Shop_Buy `json:"buys"`
	Invest int        `json:"inverst"`
}
type Shop_Buy struct {
	Cost       int    `json:"cost"`
	Collection string `json:"collection"`
}
type Recruits struct {
	Name     string `json:"name"`
	Upgraded bool   `json:"upgraded"`
}
type Roguelike_Report struct {
	Theme       int        `json:"theme"`
	Mode        string     `json:"mode"`
	Collections []string   `json:"collections"`
	Brief       Brief      `json:"brief"`
	Recruits    []Recruits `json:"recruits"`
	Initial     Initial    `json:"initial"`
	Zones       []Zone     `json:"zones"`
}
type Roguelike_Report_With_ID struct {
	ID               string `bson:"_id" json:"id"`
	Roguelike_Report `bson:",inline"`
}
type Roguelike_Report_With_UUID struct {
	UUID             interface{} `json:"uuid"`
	Roguelike_Report `bson:",inline"`
}

type ReportID struct {
	ID string `bson:"_id" json:"id" example:"6213912d7f21c24ec55377ac"`
}
type ReportIDs struct {
	ID []string `bson:"_id" json:"id" example:"6213912d7f21c24ec55377ac,6213912d7f21c24ec55377aa"`
}

/*
	About “composite literal uses unkeyed fields”：
	https://stackoverflow.com/questions/54548441/composite-literal-uses-unkeyed-fields
	Here, the Coding Writing based on MongoDB official document is used.
*/

func AddReport(r Roguelike_Report, uuid uuid.UUID) (string, error) {
	report := Roguelike_Report_With_UUID{uuid.String(), r}
	opts := options.Update().SetUpsert(true)
	ctx, cancelctx := context.WithTimeout(context.Background(), config.SERVER.READ_TIMEOUT*time.Second)
	defer cancelctx()
	result, err := db.Database("is_log").Collection("log").UpdateOne(ctx, report, bson.D{{"$set", report}}, opts)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if result.MatchedCount != 0 {
		return "", errors.New("report already exists")
	}
	log.Println("Report inserted with ID:\"", result.UpsertedID, "\"")
	return result.UpsertedID.(primitive.ObjectID).Hex(), nil
}

func AddReports(r []Roguelike_Report, uuid uuid.UUID) ([]string, error) {
	var objectids []string
	for _, d := range r {
		result, err := AddReport(d, uuid)
		if err != nil && err.Error() == "report already exists" {
			break
		} else if err != nil {
			return objectids, err
		}
		objectids = append(objectids, result)
	}
	if len(objectids) == 0 {
		return objectids, errors.New("all reports already exists")
	}
	return objectids, nil
}

func GetReportByID(logid string, result *Roguelike_Report_With_ID) error {
	objectId, err := primitive.ObjectIDFromHex(logid)
	if err != nil {
		log.Println("Invalid id\"", objectId, "\"")
		return err
	}
	filter := bson.D{{"_id", objectId}}
	ctx, cancelctx := context.WithTimeout(context.Background(), config.SERVER.READ_TIMEOUT*time.Second)
	defer cancelctx()
	res := db.Database("is_log").Collection("log").FindOne(ctx, filter)
	err = res.Decode(result)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Report\"", objectId, "\"has been found.")
	return nil
}

func GetReportListByAccount(uuid string, page int, pagesize int, result *[]Roguelike_Report_With_ID) error {
	filter := bson.D{{"uuid", uuid}}
	ctx, cancelctx := context.WithTimeout(context.Background(), config.SERVER.READ_TIMEOUT*time.Second)
	defer cancelctx()
	cursor, err := db.Database("is_log").Collection("log").Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return err
	}
	if page <= 0 || pagesize <= 0 {
		return errors.New("Page and Pagesize must larger than 0.")
	}
	for i := 1; cursor.Next(context.TODO()); i++ {
		if i > (page-1)*pagesize {
			var res Roguelike_Report_With_ID
			if err := cursor.Decode(&res); err != nil {
				log.Println(err)
				return err
			}
			*result = append(*result, res)
		}
		if i > page*pagesize {
			break
		}
		if err = cursor.Err(); err != nil {
			log.Println(err)
			return err
		}
	}
	if len(*result) == 0 {
		return errors.New("No reports was found.")
	}
	return nil
}

func DeleteReport(logid string) error {
	objectId, err := primitive.ObjectIDFromHex(logid)
	if err != nil {
		log.Println("Invalid id\"", objectId, "\"")
		return err
	}
	filter := bson.D{{"_id", objectId}}
	_, err = db.Database("is_log").Collection("log").DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Report\"", objectId, "\"has been deleted.")
	return nil
}
