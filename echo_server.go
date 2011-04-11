package main

import (
	"fmt";
	"net";
	"os";
)

type tcpPacket struct {
     data string;
     con net.Conn;
}

func handle(packet tcpPacket) {
     // just echo
     packet.con.Write([]byte(packet.data));
     //fmt.Println("Packet: ", packet);
     fmt.Printf("From socket: %w  data: '%s'\n", packet.con, packet.data);
}

func main() {
     fmt.Println("Initiating server... (Ctrl-C to stop)");
     var (
	 host = "0.0.0.0";
	 port = "4321";
	 laddr = host + ":" + port;
	 )

     data_ch := make(chan tcpPacket);
     go listen(data_ch, laddr);
     for {
	 go handle( <- data_ch );
     }
}

func listen(data_ch chan tcpPacket, laddr string) {
     fmt.Println("Listen started");
     lis, error := net.Listen("tcp", laddr);
     defer lis.Close();
     if error != nil {
	fmt.Printf("Error creating listener: %s\n", error );
	os.Exit(1);
     }
     for {
	 con, error := lis.Accept();
	 if error != nil { fmt.Printf("Error: Accepting data: %s\n", error); os.Exit(2); }
	 go listen_socket(data_ch, con);
     }
}

func listen_socket(data_ch chan tcpPacket, con net.Conn) {
     fmt.Printf("=== New Connection received from: %s \n", con.RemoteAddr());
     data   := make([]byte, 1024);
     defer con.Close();
     for {
	 n, error := con.Read(data);
	 switch error {
	      case nil:
		   packet := tcpPacket{ string(data[0:n]), con };
		   data_ch <- packet;
	      case os.EOF:
	           fmt.Printf("Warning: End of data reached: %s \n", error);
	           return;
	      default:
		   fmt.Printf("Error: Reading data : %s \n", error);
		   return;
         }
     }
}
