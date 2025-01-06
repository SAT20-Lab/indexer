package indexer

import (
	"fmt"
	"strings"

	"github.com/sat20-labs/indexer/common"
)

// 检查一个tick中的哪些nft已经被分拆
func (b *IndexerMgr) GetSplittedInscriptionsWithTick(tickerName string) []string {
	return b.ftIndexer.GetSplittedInscriptionsWithTick(tickerName)
}

func (b *IndexerMgr) GetMintPermissionInfo(ticker, address string) int64 {
	ticker = strings.ToLower(ticker)
	return b.getMintAmount(ticker, b.GetAddressId(address))
}

func (b *IndexerMgr) GetTickerMap() (map[string]*common.Ticker, error) {
	return b.ftIndexer.GetTickerMap()
}

func (b *IndexerMgr) GetOrdxTickerMapV2() (map[string]*common.TickerInfo) {
	result := make(map[string]*common.TickerInfo)
	tickers := b.ftIndexer.GetAllTickers()
	for _, tickerName := range tickers {
		t := b.GetTickerV2(tickerName)
		if t != nil {
			result[tickerName] = t
		}
	}
	return result
}

func (b *IndexerMgr) GetTicker(ticker string) *common.Ticker {
	return b.ftIndexer.GetTicker(ticker)
}

func (p *IndexerMgr) GetTickerV2(tickerName string) *common.TickerInfo {
	ticker := p.GetTicker(tickerName)
	if ticker == nil {
		return nil
	}
	result := &common.TickerInfo{}
	result.Protocol = common.PROTOCOL_NAME_ORDX
	result.Type = common.ASSET_TYPE_FT
	result.Ticker = ticker.Name
	result.DisplayName = ticker.Name
	result.Id = ticker.Id
	result.Divisibility = 0
	result.StartBlock = ticker.BlockStart
	result.EndBlock = ticker.BlockEnd
	minted, ms := p.GetMintAmount(tickerName)
	result.TotalMinted = fmt.Sprintf("%d", minted)
	result.MintTimes = ms
	result.Limit = fmt.Sprintf("%d", ticker.Limit)
	result.MaxSupply = fmt.Sprintf("%d", ticker.Max)
	result.SelfMint = ticker.SelfMint
	result.DeployHeight = int(ticker.Base.BlockHeight)
	result.DeployBlocktime = ticker.Base.BlockTime
	result.DeployTx = ""
	holders := p.GetHoldersWithTick(ticker.Name)
	result.HoldersCount = len(holders)
	result.InscriptionId = ticker.Base.InscriptionId
	result.InscriptionNum = ticker.Base.Id
	result.Description = ticker.Desc
	result.Rarity = ticker.Attr.Rarity
	result.DeployAddress, _ = p.rpcService.GetAddressByID(ticker.Base.InscriptionAddress)
	result.Content = ticker.Base.Content
	result.ContentType = string(ticker.Base.ContentType)
	result.Delegate = ticker.Base.Delegate
	return result
}

func (b *IndexerMgr) GetMintAmount(tickerName string) (int64, int64) {
	return b.ftIndexer.GetMintAmount(tickerName)
}

func (b *IndexerMgr) GetOrdxDBVer() string {
	return b.ftIndexer.GetDBVersion()
}

func (p *IndexerMgr) GetFTMintHistoryWithAddress(addressId uint64, ticker string, start int, limit int) ([]*common.InscribeBaseContent, int) {
	result := make([]*common.InscribeBaseContent, 0)
	infos, total := p.ftIndexer.GetMintHistoryWithAddress(addressId, ticker, start, limit)
	for _, info := range infos {
		mint := p.ftIndexer.GetMint(info.InscriptionId)
		if mint != nil {
			result = append(result, mint.Base)
		}
	}
	return result, total
}
