package cmd

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/ZSLTChenXiYin/game-xjson/internal/conf"
	"github.com/ZSLTChenXiYin/game-xjson/internal/converter"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

var (
	convert_cmd_config       string
	convert_cmd_table        string
	convert_cmd_output       string
	convert_cmd_sheet        string
	convert_cmd_types        bool
	convert_cmd_descriptions bool
	convert_cmd_pretty       bool
	convert_cmd_increment_id bool
	convert_cmd              = &cobra.Command{
		Use:   "convert",
		Short: "将指定xlsx数据文件转化成json数据文件",
		Long: `将指定xlsx数据文件转化成json数据文件

推荐使用方法：
  game-xjson convert -c convert.json
		`,
		Args: func(cmd *cobra.Command, args []string) error {
			if convert_cmd_config == "" && convert_cmd_table == "" {
				return errors.New("请指定配置文件或数据表")
			}

			if convert_cmd_config != "" && convert_cmd_table != "" {
				return errors.New("配置文件或数据表不能同时指定")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var convert_config *conf.ConvertConfig
			var err error

			if convert_cmd_config != "" {
				convert_config, err = conf.LoadConvertConfig(convert_cmd_config)
				if err != nil {
					return err
				}
			} else {
				convert_config = &conf.ConvertConfig{
					Table:        convert_cmd_table,
					Output:       &convert_cmd_output,
					Sheet:        &convert_cmd_sheet,
					Types:        &convert_cmd_types,
					Descriptions: &convert_cmd_descriptions,
					Pretty:       &convert_cmd_pretty,
					IncrementID:  &convert_cmd_increment_id,
				}
			}

			excel_file, err := excelize.OpenFile(convert_config.Table)
			if err != nil {
				return err
			}
			defer excel_file.Close()

			json_data, err := converter.ConvertXLSXToJSON(
				excel_file,
				*convert_config.Sheet,
				*convert_config.Types,
				*convert_config.Descriptions,
				*convert_config.IncrementID,
			)
			if err != nil {
				return err
			}

			output_file, err := os.OpenFile(*convert_config.Output, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return err
			}
			defer output_file.Close()

			if *convert_config.Pretty {
				var json_map []map[string]any
				err := json.Unmarshal(json_data, &json_map)
				if err != nil {
					return err
				}

				json_data, err = json.MarshalIndent(json_map, "", "  ")
				if err != nil {
					return err
				}

				_, err = output_file.Write(json_data)
				if err != nil {
					return err
				}
			} else {
				_, err := output_file.Write(json_data)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
)

func init() {
	root_cmd.AddCommand(convert_cmd)

	convert_cmd_flags := convert_cmd.Flags()

	convert_cmd_flags.StringVarP(&convert_cmd_config, "config", "c", "", "指定配置文件（仅json类型）")
	convert_cmd_flags.StringVarP(&convert_cmd_table, "table", "t", "", "指定xlsx数据表名")
	convert_cmd_flags.StringVarP(&convert_cmd_output, "output", "o", "output.json", "指定输出文件名")
	convert_cmd_flags.StringVarP(&convert_cmd_sheet, "sheet", "s", "", "指定xlsx工作表名（未指定时会遍历所有工作表）")
	convert_cmd_flags.BoolVarP(&convert_cmd_types, "types", "y", false, "是否指定表中的字段数据类型")
	convert_cmd_flags.BoolVarP(&convert_cmd_descriptions, "descriptions", "d", false, "是否指定表中的字段描述")
	convert_cmd_flags.BoolVarP(&convert_cmd_pretty, "pretty", "p", false, "是否格式化输出")
	convert_cmd_flags.BoolVarP(&convert_cmd_increment_id, "increment-id", "i", false, "是否在输出的json中添加自增ID（ID为计算偏移后的索引值，从1开始）")
}
