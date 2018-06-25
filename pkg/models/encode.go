// Copyright 2016 CodisLabs. All Rights Reserved.
// Licensed under the MIT (MIT-LICENSE.txt) license.

package models

import (
	"encoding/json"

	"github.com/feixiao/codis/pkg/utils/errors"
	"github.com/feixiao/codis/pkg/utils/log"
)

func jsonEncode(v interface{}) []byte {
	// MarshalIndent函数的功能和 Marshal一致，只是格式化 json，方便人工阅读。
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		log.PanicErrorf(err, "encode to json failed")
	}
	return b
}

func jsonDecode(v interface{}, b []byte) error {
	if err := json.Unmarshal(b, v); err != nil {
		return errors.Trace(err)
	}
	return nil
}
