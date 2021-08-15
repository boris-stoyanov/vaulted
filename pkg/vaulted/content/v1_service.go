// Copyright 2018 SumUp Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package content

import (
	"github.com/palantir/stacktrace"

	"github.com/sumup-oss/vaulted/pkg/vaulted/passphrase"
)

type V1Service struct {
	*baseEncryptedContentService
	aesService aesService
}

func NewV1Service(base64Service base64Service, aesService aesService) *V1Service {
	return &V1Service{
		baseEncryptedContentService: &baseEncryptedContentService{
			base64Service: base64Service,
		},
		aesService: aesService,
	}
}

func (s *V1Service) Encrypt(
	passphrase *passphrase.Passphrase,
	content *Content,
) (*EncryptedContent, error) {
	ciphertext, err := s.aesService.EncryptGCM(
		passphrase.Content,
		content.Plaintext,
	)
	if err != nil {
		return nil, stacktrace.Propagate(err, "encryption of content failed")
	}

	encryptedContent := NewEncryptedContent(ciphertext)

	return encryptedContent, nil
}

func (s *V1Service) Decrypt(
	passphrase *passphrase.Passphrase,
	encryptedContent *EncryptedContent,
) (*Content, error) {
	plaintext, err := s.aesService.DecryptGCM(
		passphrase.Content,
		encryptedContent.Ciphertext,
	)
	if err != nil {
		return nil, stacktrace.Propagate(err, "decryption of encrypted content failed")
	}

	content := NewContent(plaintext)

	return content, nil
}
