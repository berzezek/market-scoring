package utils

import (
	"market-scoring/src/config"
	"market-scoring/src/proto"
)

// Функции проверки соответствия gRPC ответа критериям пакетов

// matchesPackageXS проверяет соответствие пакету XS
func MatchesPackageXS(grpcRes *proto.DataResponse, criteria config.PackageXS) bool {
	return grpcRes.ActiveProductsAge >= int32(criteria.ActiveProductsAge)
}

// matchesPackageS проверяет соответствие пакету S
func MatchesPackageS(grpcRes *proto.DataResponse, criteria config.PackageS) bool {
	return grpcRes.ActiveProductsAge >= int32(criteria.ActiveProductsAge) &&
		grpcRes.OrderCancellationBeforeConfirmation <= int32(criteria.OrderCancellationBeforeConfirmation) &&
		grpcRes.OrderCancellationDuringDelivery <= int32(criteria.OrderCancellationDuringDelivery) &&
		grpcRes.OrderConfirmationTimeViolation <= int32(criteria.OrderConfirmationTimeViolation) &&
		grpcRes.OrderAssemblyTimeViolation <= int32(criteria.OrderAssemblyTimeViolation) &&
		grpcRes.DeliveryTimeViolation <= int32(criteria.DeliveryTimeViolation) &&
		grpcRes.SellerRating >= criteria.SellerRating
}

// matchesPackageM проверяет соответствие пакету M
func MatchesPackageM(grpcRes *proto.DataResponse, criteria config.PackageM) bool {
	return MatchesPackageS(grpcRes, criteria.PackageS)
}

// matchesPackageL проверяет соответствие пакету L
func MatchesPackageL(grpcRes *proto.DataResponse, criteria config.PackageL) bool {
	return MatchesPackageS(grpcRes, criteria.PackageS) &&
		grpcRes.Turnover >= int32(criteria.Turnover) &&
		grpcRes.IsCreditBuyButton == criteria.IsCreditBuyButton &&
		grpcRes.IsQRCreditOffline == criteria.IsQRCreditOffline
}

// matchesPackageXL проверяет соответствие пакету XL
func MatchesPackageXL(grpcRes *proto.DataResponse, criteria config.PackageXL) bool {
	return MatchesPackageL(grpcRes, criteria.PackageL)
}
