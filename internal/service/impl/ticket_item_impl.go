package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/LeVanHieu0509/backend-go/global"
	"github.com/LeVanHieu0509/backend-go/internal/consts"
	"github.com/LeVanHieu0509/backend-go/internal/database"
	"github.com/LeVanHieu0509/backend-go/internal/model"
	"github.com/LeVanHieu0509/backend-go/internal/model/mapper"
	"github.com/LeVanHieu0509/backend-go/internal/service"
	"github.com/LeVanHieu0509/backend-go/pkg/response"
)

type sTicketItem struct {
	// implementation interface here
	r                *database.Queries
	distributedCache service.IRedisCache
	localCache       service.ILocalCache
}

func NewTicketItemImpl(r *database.Queries, redisCache service.IRedisCache, localCache service.ILocalCache) *sTicketItem {
	return &sTicketItem{
		r:                r,
		distributedCache: redisCache,
		localCache:       localCache,
	}
}

func (s *sTicketItem) GetTicketItemById(ctx context.Context, ticketId int) (out model.TicketItemsOutput, err error) {

	// get data dfrom database
	// ticketItem, err := s.r.GetTicketItemById(ctx, int64(ticketId))
	// if err != nil {
	// 	return out, err
	// }
	// // mapper

	// return &model.TicketItemsOutput{
	// 	TicketId:       int(ticketItem.ID),
	// 	TicketName:     ticketItem.Name,
	// 	StockAvailable: int(ticketItem.StockAvailable),
	// 	StockInitial:   int(ticketItem.StockInitial),
	// }, nil

	// 1. get ticket item from local cache
	fmt.Println("START GET TICKET >>>>>> WITH TICKETID -> | ", ticketId)

	out, err = s.getTicketItemFromLocalCache(ctx, ticketId, "v1") // version...
	if err != nil {
		return out, fmt.Errorf("%w with id = %d -> err: %w", response.CouldNotGetTicketErr, ticketId, err)
	}

	if (out != model.TicketItemsOutput{}) {
		fmt.Println("12 - RESPONSE TICKET ITEM LOCAL CACHE -> CHECK DATA TICKET WITH ID -> ", ticketId)
		return out, nil
	}
	// 1 get cache from distributed cache
	out, err = s.getTicketItemFromDistributedCache(ctx, ticketId)
	if err != nil {
		return out, fmt.Errorf("%w with id = %d -> err: %w", response.CouldNotGetTicketErr, ticketId, err)
	}

	if (out != model.TicketItemsOutput{}) {
		fmt.Println("13 - RESPONSE TICKET ITEM DISTRIBUTED CACHE -> CHECK DATA TICKET WITH ID -> ", ticketId)
		return out, nil
	}

	//2. Nếu trong distribute cache không có data thì mới query vào database để get data
	out, err = s.getTicketItemFromDatabase(ctx, ticketId)
	if err != nil {
		return out, fmt.Errorf("%w with id = %d -> err: %w", response.CouldNotGetTicketErr, ticketId, err)
	}
	fmt.Println("11 -RESPONSE TICKET ITEM MYSQL -> CHECK DATA TICKET WITH ID -> ", ticketId)
	return out, nil
}

// get data from database
func (s *sTicketItem) getTicketItemFromDatabase(ctx context.Context, ticketId int) (out model.TicketItemsOutput, err error) {

	fmt.Println("07 - QUERY DATABASE -> CHECK DATA TICKET WITH ID -> ", ticketId)

	ticketItem, err := s.r.GetTicketItemById(ctx, int64(ticketId))
	if err != nil {
		return out, err
	}
	// add to redis cache -> not LOCK ...
	fmt.Println("08 - QUERY DATABASE: FOUND -> CHECK DATA TICKET WITH ID -> ", ticketId)
	ticketItemCacheJSON, err := json.Marshal(ticketItem)

	if err != nil {
		return out, fmt.Errorf("convert to json failed: %v", err)
	}

	err = global.Rdb.Set(
		ctx, s.getKeyTicketItemCache(ticketId),
		ticketItemCacheJSON,
		time.Duration(consts.TIME_2FA_OTP_REGISTER)*time.Minute,
	).Err()

	if err != nil {
		return out, fmt.Errorf("save redis failed: %v", err)
	}
	fmt.Println("09 - QUERY DATABASE: SET DATA TO DISTRIBUTED CACHE -> SUCCESS -> ", ticketId)

	// ticketItemCache, err := json.Marshal(ticketItemCacheJSON) // conver to byte
	// fmt.Println("ticketItemticketItem", ticketItemCacheJSON)

	isSuccess := s.localCache.SetWithTTL(ctx, s.getKeyTicketItemCache(ticketId), ticketItem)

	if !isSuccess {
		return out, fmt.Errorf("save localcache failed: %w", err)
	}

	fmt.Println("10 - QUERY DATABASE: SET DATA TO LOCAL CACHE -> SUCCESS -> ", ticketId)
	// reponse to client request
	out = mapper.ToTicketItemDTO(ticketItem) // Ở đây có mùi của java kakak...

	return out, nil
}

// get data from redis distributed
// trong hàm này có thể set lại cho localcache nếu như redis còn dữ liệu.
func (s *sTicketItem) getTicketItemFromDistributedCache(ctx context.Context, ticketId int) (out model.TicketItemsOutput, err error) {

	// ticketItemCache, err := global.Rdb.Get(ctx, s.getKeyTicketItemCache(ticketId)).Result()
	// if err != nil {
	// 	if errors.Is(err, redis.Nil) {
	// 		// Trả về lỗi riêng khi key không tồn tại
	// 		return out, nil
	// 	}
	// 	return out, fmt.Errorf("failed to get ticket item cache: %v", err)
	// }
	fmt.Println("04 - DISTRIBUTED CACHE -> CHECK DATA TICKET WITH ID -> ", ticketId)

	ticketItemCache, err := s.distributedCache.Get(ctx, s.getKeyTicketItemCache(ticketId))
	if err != nil {
		return out, fmt.Errorf("failed to get ticket item cache: %v", err)
	}
	if ticketItemCache == "" {
		fmt.Println("05 - DISTRIBUTED CACHE: NOT FOUND -> CHECK DATA TICKET WITH ID -> ", ticketId)
		return out, nil
	}

	if err := json.Unmarshal([]byte(ticketItemCache), &out); err != nil {
		return out, fmt.Errorf("parse redis data failed: %v", err)
	}
	fmt.Println("06 - DISTRIBUTED CACHE: FOUND -> CHECK DATA TICKET WITH ID -> ", ticketId, ticketItemCache)
	// put to local cache
	s.localCache.SetWithTTL(ctx, s.getKeyTicketItemCache(ticketId), out)
	return out, nil
}

// get data from local cache
func (s *sTicketItem) getTicketItemFromLocalCache(ctx context.Context, ticketId int, version string) (out model.TicketItemsOutput, err error) {
	// global.Logger.Info("getTicketItemFromLocalCache with ticketId: ", zap.Int("TicketId", ticketId))
	// var out model.TicketItemsOutput
	fmt.Println("01 - LOCAL CACHE -> CHECK DATA TICKET WITH ID -> ", ticketId)

	ticketItemLocalCache, isFound := s.localCache.Get(ctx, s.getKeyTicketItemCache(ticketId))
	if !isFound {
		// global.Logger.Info("getTicketItemFromLocalCache with ticketId is notfound: ", zap.Int("TicketId", ticketId))
		// fmt.Println("getTicketItemFromLocalCache with ticketId is notfound: ", ticketId)
		fmt.Println("02 - LOCAL CACHE: NOT FOUND -> CHECK DATA TICKET WITH ID -> ", ticketId)
		return out, nil
	}
	// fmt.Println(">>>", ticketItemLocalCache)
	// fmt.Printf(">>> Value: %+v, Type: %s\n", ticketItemLocalCache, reflect.TypeOf(ticketItemLocalCache))

	fmt.Println("03 - LOCAL CACHE: FOUND -> CHECK DATA TICKET WITH ID -> ", ticketId)

	// // Type Assertion to string
	jsonTicketString, ok := ticketItemLocalCache.(string)
	if !ok {
		fmt.Printf("ERROR: Local cache item with key %d is not a string\n", ticketId)
	}

	if err := json.Unmarshal([]byte(jsonTicketString), &out); err != nil {
		return out, fmt.Errorf("parse redis data failed: %v", err)
	}

	return out, nil
}

// util
func unmarshalTicketItem(data []byte) (model.TicketItemsOutput, error) {
	var item model.TicketItemsOutput
	err := json.Unmarshal(data, &item)
	return item, err
}

// generate key cache
func (s *sTicketItem) getKeyTicketItemCache(ticketId int) string {
	return "PRO_TICKET:ITEM:" + strconv.Itoa(ticketId)
}

/*
1. Với 1 node đơn thì sẽ tầm 24k req/s với chi phí rất nhỏ.
nhược điểm làm sao để nhất quán được 3 dữ liệu thuộc local cache, distribute cache, database.

--------- Lần đầu tiên user truy cập ------------

START GET TICKET >>>>>> WITH TICKETID -> |  1
01 - LOCAL CACHE -> CHECK DATA TICKET WITH ID ->  1
02 - LOCAL CACHE: NOT FOUND -> CHECK DATA TICKET WITH ID ->  1
04 - DISTRIBUTED CACHE -> CHECK DATA TICKET WITH ID ->  1
05 - DISTRIBUTED CACHE: NOT FOUND -> CHECK DATA TICKET WITH ID ->  1
07 - QUERY DATABASE -> CHECK DATA TICKET WITH ID ->  1
08 - QUERY DATABASE: FOUND -> CHECK DATA TICKET WITH ID ->  1
09 - QUERY DATABASE: SET DATA TO DISTRIBUTED CACHE -> SUCCESS ->  1
10 - QUERY DATABASE: SET DATA TO LOCAL CACHE -> SUCCESS ->  1
11 -RESPONSE TICKET ITEM MYSQL -> CHECK DATA TICKET WITH ID ->  1

--------- Trong 5p user truy cập thì sẽ get data ở local cache ------------
START GET TICKET >>>>>> WITH TICKETID -> |  1
01 - LOCAL CACHE -> CHECK DATA TICKET WITH ID ->  1
03 - LOCAL CACHE: FOUND -> CHECK DATA TICKET WITH ID ->  1
04 - DISTRIBUTED CACHE -> CHECK DATA TICKET WITH ID ->  1
06 - DISTRIBUTED CACHE: FOUND -> CHECK DATA TICKET WITH ID ->  1 {"ID":1,"Name":"Sach","StockInitial":1,"StockAvailable":1}
07 - QUERY DATABASE -> CHECK DATA TICKET WITH ID ->  1
08 - QUERY DATABASE: FOUND -> CHECK DATA TICKET WITH ID ->  1
09 - QUERY DATABASE: SET DATA TO DISTRIBUTED CACHE -> SUCCESS ->  1
10 - QUERY DATABASE: SET DATA TO LOCAL CACHE -> SUCCESS ->  1
11 -RESPONSE TICKET ITEM MYSQL -> CHECK DATA TICKET WITH ID ->  1

--------- Quá 5p user truy cập thì sẽ get data ở distribute cache ------------

------------- Quá thời gian distribute cache thì user sẽ thực hiện lấy data ở database -------------
*/
