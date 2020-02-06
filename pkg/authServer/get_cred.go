//
package authServer

import (
	pb "RoomStatus/proto"

	"bufio"
	"log"
	"os"

	"RoomStatus/insecure"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetCred(context.Context, *CredReq) (*CreateCredResp, error)

func (CAB *CreditsAuthBackend) GetCred(cq *pb.CredReq, stream pb.CreditsAuth_GetCredServer) error {
	CAB.mu.Lock()
	defer CAB.mu.Unlock()

	var result []*UserCredMod
	CAB.DBconn.Where(&UserCredMod{
		Username: cq.Username,
	}).Find(&result)

	if len(result) != 1 {
		return status.Error(codes.Internal, "UNKNOWN_DB_RECORD")
	}

	err := bcrypt.CompareHashAndPassword([]byte(result[0].Password), []byte(cq.Password))
	if err != nil {
		return status.Error(codes.Unauthenticated, ("PASSWORD_INVALID"))
	}

	file, err := os.Open(insecure.PrivateCertFile)
	if err != nil {
		log.Fatal(err)
		return status.Error(codes.Internal, "CRED_FILE_ISSUE")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err := stream.Send(&pb.CreateCredResp{
			Code:     200,
			File:     scanner.Bytes(),
			ErrorMsg: nil,
		}); err != nil {
			return status.Error(codes.Internal, "STREAM_SEND_ERROR")
		}
		if err := scanner.Err(); err != nil {
			return status.Error(codes.Internal, "STREAM_SCAN_ERROR")
		}
	}
	log.Println("<<End Of Send>>")
	return nil
}
