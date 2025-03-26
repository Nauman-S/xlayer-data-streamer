package datastreamer

import (
	"fmt"
	"github.com/0xPolygonHermez/zkevm-data-streamer/datastream"
	"google.golang.org/protobuf/proto"
)

func UnmarshalBookMark(data []byte) (*datastream.BookMark, error) {
	bookmark := &datastream.BookMark{}
	if err := proto.Unmarshal(data, bookmark); err != nil {
		return nil, err
	}
	return bookmark, nil
}

func UnmarshalGerUpdate(data []byte) (*datastream.UpdateGER, error) {
	update := &datastream.UpdateGER{}
	if err := proto.Unmarshal(data, update); err != nil {
		return nil, err
	}
	return update, nil
}

func UnmarshalBatchStart(data []byte) (*datastream.BatchStart, error) {
	batch := &datastream.BatchStart{}
	if err := proto.Unmarshal(data, batch); err != nil {
		return nil, err
	}
	return batch, nil
}

func UnmarshalBatchEnd(data []byte) (*datastream.BatchEnd, error) {
	batch := &datastream.BatchEnd{}
	if err := proto.Unmarshal(data, batch); err != nil {
		return nil, err
	}
	return batch, nil
}

func UnmarshalL2Block(data []byte) (*datastream.L2Block, error) {
	block := &datastream.L2Block{}
	if err := proto.Unmarshal(data, block); err != nil {
		return nil, err
	}
	return block, nil
}

func UnmarshalL2BlockEnd(data []byte) (*datastream.L2BlockEnd, error) {
	block := &datastream.L2BlockEnd{}
	if err := proto.Unmarshal(data, block); err != nil {
		return nil, err
	}
	return block, nil
}

func UnmarshalTransaction(data []byte) (*datastream.Transaction, error) {
	tx := &datastream.Transaction{}
	if err := proto.Unmarshal(data, tx); err != nil {
		return nil, err
	}
	return tx, nil
}

func parseFileEntry(entryType EntryType, data []byte) (interface{}, error) {
	switch entryType {
	case EntryType(datastream.EntryType_ENTRY_TYPE_TRANSACTION):
		return UnmarshalTransaction(data)
	//case EntryType(1):
	//	return UnmarshalBookMark(data)
	case EntryType(datastream.EntryType_ENTRY_TYPE_L2_BLOCK):
		return UnmarshalL2Block(data)
	case EntryType(datastream.EntryType_ENTRY_TYPE_L2_BLOCK_END):
		return UnmarshalL2BlockEnd(data)
	case EntryType(datastream.EntryType_ENTRY_TYPE_BATCH_START):
		return UnmarshalBatchStart(data)
	case EntryType(datastream.EntryType_ENTRY_TYPE_BATCH_END):
		return UnmarshalBatchEnd(data)
	case EntryType(datastream.EntryType_ENTRY_TYPE_UPDATE_GER):
		return UnmarshalGerUpdate(data)
	default:
		return nil, fmt.Errorf("unknown entry type: %d", entryType)
	}
}

func (e EntryType) String() string {
	switch e {
	case EntryType(datastream.EntryType_ENTRY_TYPE_TRANSACTION):
		return "transaction"

	case EntryType(datastream.EntryType_ENTRY_TYPE_L2_BLOCK):
		return "l2_block"
	case EntryType(datastream.EntryType_ENTRY_TYPE_L2_BLOCK_END):
		return "l2_block_end"
	case EntryType(datastream.EntryType_ENTRY_TYPE_BATCH_START):
		return "batch_start"
	case EntryType(datastream.EntryType_ENTRY_TYPE_BATCH_END):
		return "batch_end"
	case EntryType(datastream.EntryType_ENTRY_TYPE_UPDATE_GER):
		return "update_ger"
	default:
		return "unknown"
	}
}
