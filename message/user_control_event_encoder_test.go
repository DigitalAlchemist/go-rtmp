//
// Copyright (c) 2018- yutopp (yutopp@gmail.com)
//
// Distributed under the Boost Software License, Version 1.0. (See accompanying
// file LICENSE_1_0.txt or copy at  https://www.boost.org/LICENSE_1_0.txt)
//

package message

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserControlEventEncoderCommon(t *testing.T) {
	for _, tc := range uceTestCases {
		tc := tc // capture

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			buf := new(bytes.Buffer)

			enc := NewUserControlEventEncoder(buf)
			err := enc.Encode(tc.Value)
			assert.Nil(t, err)
			assert.Equal(t, tc.Binary, buf.Bytes())
		})
	}
}
