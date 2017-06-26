package main

import (
	"context"
	tf "github.com/dreyk/tensorflow-serving-go/pkg/tensorflow/core/framework"
	"github.com/dreyk/tensorflow-serving-go/pkg/tensorflow_serving/apis"
	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	"time"
	"fmt"

)

func main() {
	conn,err := grpc.Dial("127.0.0.1:9000",grpc.WithInsecure())
	if err!=nil{
		panic(err)
	}
	client := apis.NewPredictionServiceClient(conn)
	tContext,_ := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	resp,err := client.Predict(tContext, &apis.PredictRequest{
		ModelSpec: &apis.ModelSpec{
			Name:          "mul",
			SignatureName: "add",
			Version: &google_protobuf.Int64Value{
				Value: 2,
			},
		},
		Inputs: map[string]*tf.TensorProto{"a": &tf.TensorProto{
			Dtype:tf.DataType_DT_FLOAT,
			TensorShape:&tf.TensorShapeProto{
				Dim:[]*tf.TensorShapeProto_Dim{
					{
						Size:8,
					},
				},
			},
			FloatVal:[]float32{1,1,1,1,1,1,1,0},
		}},
	})
	if err!=nil{
		panic(err)
	}
	fmt.Println(resp)
}
