package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var InstanceDB MongoInstance

func ConexionDB(uri string) error {

	if uri == "" {
		return fmt.Errorf("debes configurar la variable de entorno URI")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("error al conectar con MongoDB: %v", err)
	}

	er := client.Ping(context.Background(), nil)
	if er != nil {
		return fmt.Errorf("error al hacer ping a MongoDB: %v", er)
	} else {
		fmt.Println("Canectado a la base de datos")
	}

	var db = client.Database("Hotels")

	InstanceDB = MongoInstance{
		Client:   client,
		Database: db,
	}

	return nil
}

func DesconectarDB() {
	if InstanceDB.Client != nil {
		if err := InstanceDB.Client.Disconnect(context.Background()); err != nil {
			fmt.Printf("Error al desconectar la base de datos: %v", err)
		} else {
			fmt.Println("Desconexión exitosa de MongoDB")
		}
	}
}

func FindAll(ModelName string) []bson.M {
	model := InstanceDB.Database.Collection(ModelName)

	busqueda, err := model.Find(context.TODO(), bson.M{})

	if err != nil {
		fmt.Print("Nose pudieron encontrar los documentos")
	}
	defer busqueda.Close(context.Background())

	var result []bson.M
	if err = busqueda.All(context.Background(), &result); err != nil {
		fmt.Print("Error al buscar los documentos")
	}

	return result

}

func FindById(ModelName string, id primitive.ObjectID) bson.M {
	model := InstanceDB.Database.Collection(ModelName)

	var result bson.M

	err := model.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)

	if err != nil {
		fmt.Print("Nose pudo encontrar el documento")
		return nil
	}

	return result
}

func Create(ModelName string, obj any) string {

	now := time.Now()

	data := bson.M{}
	bytes, _ := bson.Marshal(obj)
	bson.Unmarshal(bytes, &data)

	data["created_at"] = now
	data["updated_at"] = now

	model := InstanceDB.Database.Collection(ModelName)
	_, err := model.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println(err)
		return "Error al crear el documento"
	}
	return "Documento creado exitosamente"
}

func UpdateById(ModelName string, obj any, id primitive.ObjectID) string {

	data := bson.M{}
	bytes, _ := bson.Marshal(obj)
	bson.Unmarshal(bytes, &data)

	data["updated_at"] = time.Now()

	model := InstanceDB.Database.Collection(ModelName)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": data}

	_, err := model.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return "Error al actualizar el documento"
	}
	return "Documento actualizado exitosamente"
}

func DeleteById(ModelName string, id primitive.ObjectID) string {
	model := InstanceDB.Database.Collection(ModelName)

	filter := bson.M{"_id": id}

	_, err := model.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return "Error al eliminar el documento"
	} else {
		return "Documento eliminado exitosamente"
	}
}
