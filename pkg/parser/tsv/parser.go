package tsv

import (
	"encoding/csv"
	"io"
	"os"
	"strings"

	"github.com/shynn12/biocad/internal/item"
	"github.com/shynn12/biocad/pkg/logging"
)

func Parse(file *os.File, logger *logging.Logger) (items []item.ItemDTO) {
	var item item.ItemDTO
	r := csv.NewReader(file)
	r.Comma = '\t'
	headers, err := r.Read()
	if err != nil {
		logger.Error("Empty file")
	}
	hash := make(map[string]string)
	logger.Info("scanner is good")
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		logger.Info("Line: ", line)
		for j := 0; j < len(headers); j++ {
			hash[strings.TrimSpace(strings.ToLower(headers[j]))] = strings.TrimSpace(line[j])
		}
		logger.Info("Putting a data into structure")
		item.Number = hash["n"]
		item.Mqtt = hash["mqtt"]
		item.Invid = hash["invid"]
		item.UnitGuid = hash["unit_guid"]
		item.MsgId = hash["msg_id"]
		item.Text = hash["text"]
		item.Context = hash["context"]
		item.Class = hash["class"]
		item.Level = hash["level"]
		item.Area = hash["area"]
		item.Addr = hash["addr"]
		item.Block = hash["block"]
		item.Type = hash["type"]
		item.Bit = hash["bit"]
		item.InvertBit = hash["invert_bit"]
		items = append(items, item)
	}
	return items
}
