package main

import (
	// "bytes"
	// "image/jpeg"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/bluenviron/gortsplib/v4"
	"github.com/bluenviron/gortsplib/v4/pkg/format"
	"github.com/bluenviron/gortsplib/v4/pkg/format/rtph264"
	"github.com/bluenviron/gortsplib/v4/pkg/url"
	"github.com/pion/rtp"
	"github.com/zeromq/goczmq"
)

// This example shows how to
// 1. connect to a RTSP server
// 2. check if there's an H264 media stream
// 3. decode the H264 media stream into RGBA frames

// This example requires the FFmpeg libraries, that can be installed with this command:
// apt install -y libavformat-dev libswscale-dev gcc pkg-config
func main() {

	// url_zmq := "tcp://127.0.0.1:5555"
	a := []string{
		// "rtsp://192.168.91.150:8554/80aebd2b-11c5-471f-8b30-7316ab35dadb",loi
		// "rtsp://192.168.91.150:8554/9ff82d0b-648c-41e1-88b8-efef3e513433",loi
		// "rtsp://192.168.91.150:8554/96928301-2873-417a-87b3-769759dec9f6",
		// "rtsp://192.168.91.150:8554/5e4ee879-b748-4bbb-b436-590258bf0d5f",
		// "rtsp://192.168.91.150:8554/9db76c33-2e97-41c2-b633-6b0fa3fc2a02",
		// "rtsp://192.168.91.150:8554/900abac4-793c-4b7d-8a2f-6256d5b4459c",
		// "rtsp://192.168.91.150:8554/be378ef5-e080-4e7b-b285-491225d206c8",
		// "rtsp://192.168.91.150:8554/324cd819-e22d-4c98-aaeb-4e6a3d544c22",
		// "rtsp://192.168.91.150:8554/a6bcad22-84fd-41c0-b718-fc877ff18b6c",
		// "rtsp://192.168.91.150:8554/81356508-ddf9-4b19-8413-f6e41dab2194",loi
		// "rtsp://192.168.91.150:8554/5ba8b4a9-9460-4770-b49c-3046d3a9057e",loi
		// "rtsp://192.168.91.150:8554/f258e16a-66c6-4279-b5ac-70c7f8663ca6",
		// "rtsp://192.168.91.150:8554/c2ed9aa6-8fb0-4bef-ab86-443d215487b5",
		// "rtsp://192.168.91.150:8554/44e309e2-8107-43ea-81cc-12c3b6c8350e",
		// "rtsp://192.168.91.150:8554/28818a02-aa23-4b3a-a4dd-3d7d40ba77c8",loi
		// "rtsp://192.168.91.150:8554/dd69f608-db98-4b88-846d-a3029eec72d1",loi
		// "rtsp://192.168.91.150:8554/6a0013a4-f8a1-417e-bd19-8337a70d536b",
		// "rtsp://192.168.91.150:8554/ef3e1d2a-885f-4de0-9d4e-5785a316cde2",loi
		// "rtsp://192.168.91.150:8554/73e8f87c-0c22-423a-9548-5cf0e79d4507",loi
		// "rtsp://192.168.91.150:8554/c185cccc-fa87-4d85-a031-05a84492c371",
		// "rtsp://192.168.91.150:8554/b6ddd63e-ba48-4765-8574-a99caec18f85",
		// "rtsp://192.168.91.150:8554/3c16b7f9-8d48-4d3e-8934-cf246a91ef47",loi
		// "rtsp://192.168.91.150:8554/87fdb891-d357-4a25-9d43-25d1743ca562",
		// "rtsp://192.168.91.150:8554/7d1ce41f-8091-492f-8cee-58b02070d12f",
		// "rtsp://192.168.91.150:8554/716348ca-2c45-47c3-9a77-8efca03374d0",loi
		// "rtsp://192.168.91.150:8554/c7f2c8f9-a8a7-43ad-a7d6-390a911b37c4",
		// "rtsp://192.168.91.150:8554/5795d4e3-91ca-485e-bc10-ae69b19f32b4",
		// "rtsp://192.168.91.150:8554/b525f84d-33d6-42ac-9b4a-0c121334bb58",
		// "rtsp://192.168.91.150:8554/c3256979-896d-4676-bea8-8a7fba0d32c8",
		// "rtsp://192.168.91.150:8554/772bff6d-d713-45f7-96f4-5654126f1db1",
		// "rtsp://192.168.91.150:8554/624102b3-50cb-4ed3-abe5-8d35ef042f6d",
		// "rtsp://192.168.91.150:8554/9c738fd0-31ae-4276-902c-35e52ef6f09f",loi
		// "rtsp://192.168.91.150:8554/a94bda88-90bf-47c6-a33c-cbc400746655",loi
		// "rtsp://192.168.91.150:8554/35a53150-490c-426d-9a3b-f761c10b12a1",
		// "rtsp://192.168.91.150:8554/34e7a74f-fbe6-49e8-8c4e-b98cfef4acf5",
		// "rtsp://192.168.91.150:8554/deahan",
		"rtsp://192.168.91.150:8554/deahan1",
		// "rtsp://192.168.91.150:8554/db099347-d123-430e-a3ea-7706894f05d3",
		// "rtsp://192.168.91.150:8554/e219a1f2-37b8-4f88-bc88-6296f8c6c4b2",loi
		// "rtsp://192.168.91.150:8554/7a916881-cd13-4f76-8f49-52eccb8a2766",loi
		// "rtsp://192.168.91.150:8554/7f4d03dc-8a25-4711-bd93-8d5f4f5077d6",loi
		// "rtsp://192.168.91.150:8554/03f8f97c-9843-4f6c-9cc4-27dcdc2b4eba",loi
		// "rtsp://192.168.91.150:8554/672237ad-f8ce-4bb2-aa88-c4296e14330a",loi
		// "rtsp://192.168.91.150:8554/9f4754c9-ef11-462e-8480-eb140ea06831",loi
		// "rtsp://192.168.91.150:8554/12591272-dc56-419b-8417-9725776b5419",loi
		// "rtsp://192.168.91.150:8554/0992d849-6e83-4388-aa7e-8074b6f7e334",loi
		// "rtsp://192.168.91.150:8554/74ec1bed-dc85-402e-951d-4e58508f1c54",loi
		// "rtsp://192.168.91.150:8554/8f380548-166e-46cf-8c0b-a485b4012213",loi
		// "rtsp://192.168.91.150:8554/25bb93aa-c0d6-411d-be0d-f86114748063",loi
		// "rtsp://192.168.91.150:8554/3b11bb07-eb94-41bb-bc0f-e6e01472968f",loi
		// "rtsp://192.168.91.150:8554/cf2c79f9-e23b-4d95-9bb7-892ac28f0b66",loi
		// "rtsp://192.168.91.150:8554/d5f55fec-0c11-446e-80d4-3c45cdf15cad",loi
		// "rtsp://192.168.91.150:8554/49b21433-8068-4efa-9253-89357060612b",loi
		// "rtsp://192.168.91.150:8554/701b9a54-6970-4bc0-a12f-77621da5022e",loi
		// "rtsp://192.168.91.150:8554/ac1e8cec-81bb-4a90-b693-d2bc09c8dd61",loi
		// "rtsp://192.168.91.150:8554/49316f2d-494e-4a13-ac1e-be76ee4264ab",loi
		// "rtsp://192.168.91.150:8554/7741b94b-5c1f-411c-97ec-7dfead510eac",
		// "rtsp://192.168.91.150:8554/deahan2",loi
		// "rtsp://192.168.91.150:8554/b194b225-eab6-4632-b2f6-7a784118d200",loi
		// "rtsp://192.168.91.150:8554/7ed73add-a99f-4ab8-b944-ccbe4efa7359",loi
	}
	var wg sync.WaitGroup
	wg.Add(len(a))
	base_port := 1122
	println(len(a))

	for i, s := range a {
		port := base_port + i
		go run_thread(s,port, &wg)
	}
	wg.Wait()
}

func run_thread(url_zmq string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	println(url_zmq)
	zmq_url := "tcp://192.168.91.152:" + strconv.Itoa(port)
	//zmq_url := "tcp://localhost:" + strconv.Itoa(port)
	// Create a dealer socket and connect it to the router.
	// dealer, err := goczmq.NewPub("tcp://192.168.91.150:1122")
	dealer, err := goczmq.NewPub(zmq_url)
	if err != nil {
		log.Fatal(err)
	}
	defer dealer.Destroy()

	// log.Println("dealer created and connected to: tcp://192.168.91.150:1122")
	// log.Println("start")
	// log.Println(url_zmq)
	c := gortsplib.Client{}
	c.ReadTimeout = 200*time.Second
	// parse URL
	// u, err := url.Parse("rtsp://192.168.91.150:8554/deahan1")
	u, err := url.Parse(url_zmq)
	if err != nil {
		panic(err)
	}

	// connect to the server
	err = c.Start(u.Scheme, u.Host)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// find published medias
	desc, _, err := c.Describe(u)
	if err != nil {
		panic(err)
	}

	// find the H264 media and format
	var forma *format.H264
	medi := desc.FindFormat(&forma)
	if medi == nil {
		panic("media not found")
	}

	// setup RTP/H264 -> H264 decoder
	rtpDec, err := forma.CreateDecoder()
	if err != nil {
		panic(err)
	}

	// setup H264 -> raw frames decoder
	frameDec, err := newH264Decoder()
	if err != nil {
		panic(err)
	}
	defer frameDec.close()

	// if SPS and PPS are present into the SDP, send them to the decoder
	if forma.SPS != nil {
		frameDec.decode(forma.SPS)
	}
	if forma.PPS != nil {
		frameDec.decode(forma.PPS)
	}

	// setup a single media
	_, err = c.Setup(desc.BaseURL, medi, 0, 0)
	if err != nil {
		panic(err)
	}

	// called when a RTP packet arrives
	c.OnPacketRTP(medi, forma, func(pkt *rtp.Packet) {
		// decode timestamp
		_, ok := c.PacketPTS(medi, pkt)
		if !ok {
			// log.Printf("waiting for timestamp")
			return
		}

		// extract access units from RTP packets
		au, err := rtpDec.Decode(pkt)
		// _, err := rtpDec.Decode(pkt)
		if err != nil {
			if err != rtph264.ErrNonStartingPacketAndNoPrevious && err != rtph264.ErrMorePacketsNeeded {
				log.Printf("ERR: %v", err)
			}
			return
		}

		for _, nalu := range au {
			byteSlice := make([]byte, len(nalu))
			// Convert each unint8 element to a byte element.
			for i, v := range nalu {
				byteSlice[i] = byte(v)
			//        log.Printf("v = %T\n", v)
			}
			if len(nalu) > 10 {
				log.Printf("byteSlice = %v\n", byteSlice[:10])
			}
			err = dealer.SendFrame(nalu, goczmq.FlagNone)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(zmq_url," send img bytes array with len: ", len(nalu))
		}
	})
	// start playing
	_, err = c.Play(nil)
	if err != nil {
		panic(err)
	}

	// wait until a fatal error
	// panic(c.Wait())
	err = c.Wait()
	if err != nil {
	// Some other error occurred.
		log.Fatal(err)
	}

}
