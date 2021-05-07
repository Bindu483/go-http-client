package aws

import (
	"crypto/rsa"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/reynencourt/rc-common-lib/v2/ops/sshkeygen"
)

func (a *Provider) GenerateAndUploadKeyPair(awsKeyName string) (keyname string, privateKey []byte, publicKey []byte, err error) {

	var (
		privateKeyRaw *rsa.PrivateKey
		importResp    *ec2.ImportKeyPairOutput
	)

	ec2Session := ec2.New(a.session)

	keygen := sshkeygen.New(awsKeyName, 4096)

	privateKeyRaw, err = keygen.GeneratePrivateKey()
	if err != nil {
		return
	}

	publicKey, err = keygen.GeneratePublicKey(&privateKeyRaw.PublicKey)
	if err != nil {
		return
	}

	privateKeyBytes := keygen.EncodePrivateKeyToPEM(privateKeyRaw)

	err = keygen.WriteKeyToFile(privateKeyBytes, publicKey)
	if err != nil {
		return
	}

	importResp, err = ec2Session.ImportKeyPair(&ec2.ImportKeyPairInput{
		KeyName:           aws.String(awsKeyName),
		DryRun:            aws.Bool(false),
		PublicKeyMaterial: publicKey,
	})

	if err != nil {
		return
	}
	if importResp == nil {
		err = errors.New("no response from aws")
		return
	}

	keyname = *importResp.KeyName

	return
}
