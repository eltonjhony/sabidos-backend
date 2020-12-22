package config

import (
	"context"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db *mongo.Client
var port string

var collection *mongo.Collection
var ctx = context.TODO()

func ConnectToDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://db:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected to MongoDB!")
	}

	return client
}

func SetupModels(ac entity.AccountDataProvider, av entity.AvatarDataProvider, cat entity.CategoryDataProvider, qz entity.QuizDataProvider, conn *mongo.Client) {
	quickstartDatabase := conn.Database("sabidos")
	quickstartDatabase.Collection("accounts")

	avatar1 := entity.Avatar{1, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167774/sabidos/avatar/1.png"}
	avatar2 := entity.Avatar{2, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167776/sabidos/avatar/2.png"}
	avatar3 := entity.Avatar{3, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167776/sabidos/avatar/3.png"}
	avatar4 := entity.Avatar{4, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167776/sabidos/avatar/4.png"}
	avatar5 := entity.Avatar{5, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167775/sabidos/avatar/5.png"}
	avatar6 := entity.Avatar{6, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167775/sabidos/avatar/6.png"}
	avatar7 := entity.Avatar{7, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167775/sabidos/avatar/7.png"}
	avatar8 := entity.Avatar{8, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167775/sabidos/avatar/8.png"}
	avatar9 := entity.Avatar{9, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167775/sabidos/avatar/9.png"}
	avatar10 := entity.Avatar{10, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167775/sabidos/avatar/10.png"}
	avatar11 := entity.Avatar{11, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167774/sabidos/avatar/11.png"}
	avatar12 := entity.Avatar{12, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167774/sabidos/avatar/12.png"}
	avatar13 := entity.Avatar{13, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167774/sabidos/avatar/13.png"}
	avatar14 := entity.Avatar{14, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167774/sabidos/avatar/14.png"}
	avatar15 := entity.Avatar{15, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167774/sabidos/avatar/15.png"}

	account, _ := ac.GetByIdentifier(context.Background(), "smash", "")

	if len(account.NickName) == 0 {
		newAcc := entity.Account{"yiXtigKxtEVKl5mBh4qB7ZKumBs1", "Hulk", "Smash", entity.Avatar{1, ""}, entity.Reputation{5, 10}, 3, 100, 100, "email", true, "tel"}
		ac.Insert(context.Background(), newAcc)
	}

	avatar, _ := av.FindById(context.Background(), 1)

	if avatar.Id == 0 {

		av.Insert(context.Background(), avatar1)
		av.Insert(context.Background(), avatar2)
		av.Insert(context.Background(), avatar3)
		av.Insert(context.Background(), avatar4)
		av.Insert(context.Background(), avatar5)
		av.Insert(context.Background(), avatar6)
		av.Insert(context.Background(), avatar7)
		av.Insert(context.Background(), avatar8)
		av.Insert(context.Background(), avatar9)
		av.Insert(context.Background(), avatar10)
		av.Insert(context.Background(), avatar11)
		av.Insert(context.Background(), avatar12)
		av.Insert(context.Background(), avatar13)
		av.Insert(context.Background(), avatar14)
		av.Insert(context.Background(), avatar15)

	}

	category1 := entity.Category{1, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1606393715/sabidos/categories/aleatorio_g89mxu.jpg", "Aleatório", "https://res.cloudinary.com/ddb86uj5i/image/upload/v1602365131/sabidos/categories/ic_random_category_hvxa1a.png"}
	category2 := entity.Category{2, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1606394231/sabidos/categories/curiosidades_fiafjk.jpg", "Curiosidades", "https://res.cloudinary.com/ddb86uj5i/image/upload/v1602367579/sabidos/categories/ic_curiosity_category_hjmqrl.png"}
	category3 := entity.Category{3, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1606394884/sabidos/categories/esportes_aoasbl.jpg", "Esportes", "https://res.cloudinary.com/ddb86uj5i/image/upload/v1602367765/sabidos/categories/ic_sports_category_negv0m.png"}
	category4 := entity.Category{4, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1606393149/sabidos/categories/ciencias_qvdj7i.jpg", "Ciências", "https://res.cloudinary.com/ddb86uj5i/image/upload/v1602367917/sabidos/categories/ic_science_category_qbtqsn.png"}
	category5 := entity.Category{5, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1606393149/sabidos/categories/geografia_af5cnr.jpg", "Geografia", "https://res.cloudinary.com/ddb86uj5i/image/upload/v1602368011/sabidos/categories/ic_geography_category_urnzwm.png"}
	category6 := entity.Category{6, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1606393149/sabidos/categories/historia_o6bnot.jpg", "História", "https://res.cloudinary.com/ddb86uj5i/image/upload/v1602368097/sabidos/categories/ic_history_category_myoxdd.png"}
	category7 := entity.Category{7, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1606391498/sabidos/categories/biologia_oakxlb.jpg", "Biologia", "https://res.cloudinary.com/ddb86uj5i/image/upload/v1602372897/sabidos/categories/ic_bio_category_yzm1li.png"}
	category8 := entity.Category{8, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1606395384/sabidos/categories/literatura_y8a5qe.jpg", "Literatura", "https://res.cloudinary.com/ddb86uj5i/image/upload/v1602372999/sabidos/categories/ic_literatura_category_me8qkt.png"}
	category9 := entity.Category{9, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1606395384/sabidos/categories/religiao_vfflhm.jpg", "Religião", "https://res.cloudinary.com/ddb86uj5i/image/upload/v1602373099/sabidos/categories/ic_religiao_category_se2p2j.png"}
	category10 := entity.Category{10, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1606387869/sabidos/categories/cinema_lq0mqm.jpg", "Cinema", "https://res.cloudinary.com/ddb86uj5i/image/upload/v1602373303/sabidos/categories/ic_cine_category_jgzjet.png"}
	category11 := entity.Category{11, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1606391087/sabidos/categories/tecnologia_w8skxx.jpg", "Tecnologia", "https://res.cloudinary.com/ddb86uj5i/image/upload/v1602373389/sabidos/categories/ic_tech_category_dqfmsj.png"}
	category12 := entity.Category{12, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1606390699/sabidos/categories/games_yixylw.jpg", "Games", "https://res.cloudinary.com/ddb86uj5i/image/upload/v1602373462/sabidos/categories/ic_games_category_fdhior.png"}

	category, _ := cat.FindById(context.Background(), 1)

	if category.Id == 0 {

		cat.Insert(context.Background(), category1)
		cat.Insert(context.Background(), category2)
		cat.Insert(context.Background(), category3)
		cat.Insert(context.Background(), category4)
		cat.Insert(context.Background(), category5)
		cat.Insert(context.Background(), category6)
		cat.Insert(context.Background(), category7)
		cat.Insert(context.Background(), category8)
		cat.Insert(context.Background(), category9)
		cat.Insert(context.Background(), category10)
		cat.Insert(context.Background(), category11)
		cat.Insert(context.Background(), category12)

	}

	quiz1 := entity.Quiz{
		ImageUrl:           "https://res.cloudinary.com/ddb86uj5i/image/upload/v1602710672/sabidos/quiz/turtle_jgcttl.jpg",
		Description:        "Qual ciência estuda especificamente a vida animal e vegetal no mar?",
		QuizLimitInSeconds: 15,
		Category:           category4,
		Alternatives:       []entity.Alternative{},
		Explanation: entity.Explanation{
			Description: "Biologia marinha é a especialidade biológica que se encarrega de entender os organismos que vivem em ecossistemas de água salgada e as relações dos mesmos com o ambiente. Os oceanos cobrem mais de 71% da superfície da Terra e, assim como o ambiente terrestre é diverso, os oceanos também são. Por isso encontramos as mais diferentes formas de vida no mar, desde o plâncton microscópico, incluindo o fitoplâncton, de enorme importância para a produção primária.",
			Resource:    "Wikipédia.org",
		},
	}
	quiz1.AddAlternative(entity.Alternative{
		Description:        "Ictiologia",
		IsCorrect:          false,
		PercentageAnswered: 10,
	})
	quiz1.AddAlternative(entity.Alternative{
		Description:        "Biologia Marinha",
		IsCorrect:          true,
		PercentageAnswered: 80,
	})
	quiz1.AddAlternative(entity.Alternative{
		Description:        "Bioquímica",
		IsCorrect:          false,
		PercentageAnswered: 3,
	})
	quiz1.AddAlternative(entity.Alternative{
		Description:        "Hipertologia",
		IsCorrect:          false,
		PercentageAnswered: 7,
	})

	qz.Insert(context.Background(), quiz1)

}

func SetUpDBConnection(DB *mongo.Client) {
	db = DB
}

func GetDBConnection() *mongo.Client {
	return db
}

func SetPortConnection(Port string) {
	port = Port
}

func GetPortConnection() string {
	return port
}
