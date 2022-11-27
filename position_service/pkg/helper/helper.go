package helper

import (
	"database/sql"
	"time"

	"bitbucket.org/Udevs/position_service/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleError(log logger.Logger, err error, message string, req interface{}, code codes.Code) error {
	if code == codes.Canceled {
		log.Error(message, logger.Error(err), logger.Any("req", req))
		return status.Error(code, message)
	} else if err == sql.ErrNoRows {
		log.Error(message+", Not Found", logger.Error(err), logger.Any("req", req))
		return status.Error(codes.NotFound, "Not Found")
	} else if err != nil {
		log.Error(message, logger.Error(err), logger.Any("req", req))
		return status.Error(codes.Internal, message)
	}
	return nil
}

func ParseTime(timeString string) (time.Time, error) {
	resp, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err != nil {
		return time.Time{}, err
	}

	return resp, err
}
