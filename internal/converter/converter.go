package converter

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func ConvertXLSXToJSON(excel_file *excelize.File, sheet_name string, types bool, descriptions bool, pretty bool, increment_id bool) ([]byte, error) {
	// 检测excel文件对象
	if excel_file == nil {
		return nil, errors.New("excel文件对象不可为空")
	}

	// 检测sheet名称
	sheet_list := excel_file.GetSheetList()
	if len(sheet_list) == 0 {
		return nil, errors.New("excel文件没有sheet")
	}
	if sheet_name != "" {
		sheet_exist := false
		for _, sheet := range sheet_list {
			if sheet == sheet_name {
				sheet_exist = true
				break
			}
		}
		if !sheet_exist {
			return nil, errors.New("excel文件没有名为" + sheet_name + "的sheet")
		}
		// 如果sheet名称不为空，指定单一sheet名称
		sheet_list = []string{sheet_name}
	}

	// 创建xlsx json数据
	xlsx_json_maps := make([]map[string]any, len(sheet_list))

	// 遍历sheet列表
	for index_0, sheet := range sheet_list {
		// 创建xlsx sheet json数据
		xlsx_json_maps[index_0] = make(map[string]any)

		// 获取sheet数据
		rows, err := excel_file.GetRows(sheet)
		if err != nil {
			return nil, err
		}

		// 检测sheet数据
		if len(rows) == 0 {
			return nil, errors.New("sheet为空")
		}

		var field_names []string
		var field_types []ColumnType
		var sheet_json_maps []map[string]any
		offset := 1

		// 获取字段名称
		field_names = make([]string, len(rows[0]))
		copy(field_names, rows[0])

		// 获取字段类型并计算行偏移
		if types || descriptions {
			// 如果存在类型列
			if types {
				// 获取字段类型
				field_types = make([]ColumnType, len(rows[1]))

				for index, type_str := range rows[1] {
					column_type_exist := false

					for key, value := range type_map {
						if type_str == key {
							field_types[index] = value
							column_type_exist = true
							break
						}
					}

					if !column_type_exist {
						return nil, fmt.Errorf("sheet第2行第%d列类型错误", index+1)
					}
				}
			}

			// 计算行偏移
			if types && descriptions {
				offset = 3
			} else if types {
				offset = 2
			} else if descriptions {
				offset = 2
			}
		}

		// 创建sheet json数据
		sheet_json_maps = make([]map[string]any, len(rows)-offset)
		// 遍历行
		for index_1, row := range rows[offset:] {
			sheet_json_maps[index_1] = make(map[string]any)

			// 添加id字段
			if increment_id {
				sheet_json_maps[index_1]["id"] = index_1 + 1
			}

			// 添加字段
			for index_2, value := range row {
				field_name := field_names[index_2]
				field_type := field_types[index_2]

				// 如果存在类型列
				if types {
					accurate_value, err := convertAccurateValue(value, field_type)
					if err != nil {
						return nil, err
					}

					sheet_json_maps[index_1][field_name] = accurate_value
				} else {
					sheet_json_maps[index_1][field_name] = autoConvertValue(value)
				}
			}
		}

		// 添加sheet json数据
		xlsx_json_maps[index_0][sheet] = sheet_json_maps
	}

	var xlsx_json_data []byte
	var err error
	// 返回json数据
	if len(xlsx_json_maps) == 1 {
		if pretty {
			xlsx_json_data, err = json.MarshalIndent(xlsx_json_maps[0][sheet_list[0]].([]map[string]any), "", "  ")
			if err != nil {
				return nil, err
			}
		} else {
			xlsx_json_data, err = json.Marshal(xlsx_json_maps[0][sheet_list[0]].([]map[string]any))
			if err != nil {
				return nil, err
			}
		}
	} else {
		if pretty {
			xlsx_json_data, err = json.MarshalIndent(xlsx_json_maps, "", "  ")
			if err != nil {
				return nil, err
			}
		} else {
			xlsx_json_data, err = json.Marshal(xlsx_json_maps)
			if err != nil {
				return nil, err
			}
		}
	}

	return xlsx_json_data, nil
}

func autoConvertValue(value string) any {
	if value == "" {
		return nil
	}

	int_value, err := strconv.Atoi(value)
	if err == nil {
		return int_value
	}

	float_value, err := strconv.ParseFloat(value, 64)
	if err == nil {
		return float_value
	}

	bool_value, err := strconv.ParseBool(value)
	if err == nil {
		return bool_value
	}

	json_value := make(map[string]any)
	err = json.Unmarshal([]byte(value), &json_value)
	if err == nil {
		return json_value
	}

	return value
}

func convertAccurateValue(value string, column_type ColumnType) (any, error) {
	switch column_type {
	case COLUMN_TYPE_STRING:
		return value, nil
	case COLUMN_TYPE_BOOL:
		if value == "" {
			return false, nil
		}

		bool_value, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}

		return bool_value, nil
	case COLUMN_TYPE_FLOAT:
		if value == "" {
			return 0.0, nil
		}

		float_value, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}

		return float_value, nil
	case COLUMN_TYPE_INT:
		if value == "" {
			return 0, nil
		}

		int_value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}

		return int_value, nil
	case COLUMN_TYPE_JSON:
		if value == "" {
			return nil, nil
		}

		json_value := make(map[string]any)
		err := json.Unmarshal([]byte(value), &json_value)
		if err != nil {
			return nil, err
		}
		return json_value, nil
	default:
		return nil, errors.New("未知类型")
	}
}
