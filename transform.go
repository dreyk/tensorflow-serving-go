package main

import (
	"image"
	_ "image/jpeg"
	"os"
	"github.com/dreyk/tensorflow-serving-go/pkg/tensorflow_serving/apis"
	"context"
	"time"
	"fmt"
	tf "github.com/dreyk/tensorflow-serving-go/pkg/tensorflow/core/framework"
	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"flag"
	"github.com/nfnt/resize"
)

var path string
func main() {
	flag.StringVar(&path,"i","","Img")
	flag.Parse()
	ifile,err := os.Open(path)
	if err!=nil{
		panic(err)
	}
	data,_,err := image.Decode(ifile)
	if err!=nil{
		panic(err)
	}
	pdata0 := []float32{}
	data = resize.Resize(256, 256, data, resize.Lanczos3)
	for y := data.Bounds().Min.Y ; y<256;y++{
		for x := data.Bounds().Min.X ; x<256;x++{
			r,g,b,_ := data.At(x,y).RGBA()
			r8 := float32(uint8(r))
			g8 := float32(uint8(g))
			b8 := float32(uint8(b))
			pdata0 = append(pdata0,r8,g8,b8)
		}
	}
	pdata := []float32{}
	for i := 0 ; i < 1; i++{
		pdata = append(pdata,pdata0...)
	}
	conn,err := grpc.Dial("127.0.0.1:9000",grpc.WithInsecure())
	if err!=nil{
		panic(err)
	}
	client := apis.NewPredictionServiceClient(conn)
	tContext,_ := context.WithTimeout(context.Background(), time.Duration(1*time.Minute))
	resp,err := client.Predict(tContext, &apis.PredictRequest{
		ModelSpec: &apis.ModelSpec{
			Name:          "styles",
			SignatureName: "transform",
			Version: &google_protobuf.Int64Value{
				Value: 1,
			},
		},
		Inputs: map[string]*tf.TensorProto{"X_content": &tf.TensorProto{
			Dtype:tf.DataType_DT_FLOAT,
			TensorShape:&tf.TensorShapeProto{
				Dim:[]*tf.TensorShapeProto_Dim{
					{
						Size:1,
					},
					{
						//Size:int64(data.Bounds().Dy()),
						Size:256,
					},
					{
						//Size:int64(data.Bounds().Dx()),
						Size:256,
					},
					{
						Size:3,
					},
				},
			},
			FloatVal:pdata,
		}},
	})
	if err!=nil{
		panic(err)
	}
	for _,v := range resp.Outputs{
		fmt.Println(*v.TensorShape)
	}
}
