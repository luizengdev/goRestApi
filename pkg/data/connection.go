package data

import (
	"GoRestApi/pkg/config"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connection struct {
	Client *mongo.Client
	ctx    context.Context
}

// Cria uma conexão com o banco de dados Mongo
func NewMongoConnection(cfg *config.Settings) Connection {
	// Definindo nosso URI para corresponder ao que foi inicializado no arquivo docker compose
	URI := fmt.Sprintf("mongodb://%s/%s", cfg.DbHost, cfg.DbName)

	// Configurando as credenciais usadas para autenticação
	credentials := options.Credential{
		Username: cfg.DbUser,
		Password: cfg.DbPass,
	}

	// Criando uma instância de opções mongo com o URI correto e as credenciais aplicadas.
	clientOpts := options.Client().ApplyURI(URI).SetAuth(credentials)

	// Um novo contexto é criado e passado para a função mongo.Connect,
	// que cancelará a operação de conexão caso demore muito.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)

	// Pingando banco de dados para garantir uma conexão adequada
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to database.")

	return Connection{
		Client: client,
		ctx:    ctx,
	}
}

// Disconnect desconectará da instância mongo quando chamada, permite controlar a func exposta de nosso pacote.
func (c Connection) Disconnect() {
	c.Client.Disconnect(c.ctx)
}
