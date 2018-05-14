package dgraph

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/emwalker/digraph/server/common"
	"google.golang.org/grpc"
)

type Root struct {
	Query []common.Link `json:"query"`
}

func Do(ctx context.Context, query string) *api.Response {
	conn, err := grpc.Dial("0.0.0.0:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	variables := map[string]string{}
	tx := dg.NewTxn()
	defer tx.Discard(ctx)

	log.Println("sending query:\n", query)
	response, err := tx.QueryWithVars(ctx, query, variables)
	if err != nil {
		log.Fatal(err)
	}
	return response
}
