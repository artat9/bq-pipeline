package arweave

import (
	"aurora-backend/lib/functions/lib/ad"
	"aurora-backend/lib/functions/lib/common/log"
	"aurora-backend/lib/functions/lib/klimg"
	"aurora-backend/lib/functions/lib/post"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Dev43/arweave-go/api"
	"github.com/Dev43/arweave-go/transactor"
	"github.com/Dev43/arweave-go/wallet"
)

const (
	uri       = "https://arweave.net"
	walletKey = "./assets/arweave/arweave.json"
)

type (
	// Client client
	Client struct {
		svc *api.Client
		w   *wallet.Wallet
	}
)

// NewClient create new client
func NewClient() (*Client, error) {
	cli, err := newClient()
	if err != nil {
		log.Error("client create failed", err)
		return nil, err
	}
	w, err := newWallet()
	if err != nil {
		log.Error("wallet import failed", err)
		return nil, err
	}

	return &Client{
		svc: cli,
		w:   w,
	}, nil
}

func newWallet() (*wallet.Wallet, error) {
	w := wallet.NewWallet()
	if err := w.LoadKeyFromFile(walletKey); err != nil {
		log.Error("load keyfile failed", err)
		return nil, err
	}
	return w, nil
}

func newClient() (*api.Client, error) {
	return api.Dial(uri)
}
func (c *Client) uploadImage(ctx context.Context, img klimg.Image) (string, error) {
	return c.sendTransaction(ctx, img.Data, img.ContentType)
}

func newTransact() (*transactor.Transactor, error) {
	return transactor.NewTransactor(uri)
}

// UploadImage upload image
func (c *Client) UploadImage(ctx context.Context, img klimg.Image) (string, error) {
	return c.sendTransaction(ctx, img.Data, img.ContentType)
}

func (c *Client) dl(ctx context.Context, txID string) ([]byte, error) {
	url := uri + "/" + txID
	res, err := http.Get(url)
	if err != nil {
		log.Error("http.Get -> %v", err)
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Error("ioutil.ReadAll -> %v", err)
		return nil, err
	}
	return data, err
}

func (c *Client) downloadImage(ctx context.Context, txID string) error {
	res, err := c.dl(ctx, txID)
	if err != nil {
		return err
	}
	ioutil.WriteFile("/tmp/project.png", res, 0666)

	if err != nil {
		log.Error("write file failed", err)
		return err
	}
	return err
}

func (c *Client) download(ctx context.Context, txID string) (string, error) {
	log.Info("txID:" + txID)
	return c.downloadWithRetry(ctx, txID, 0)
}

func (c *Client) downloadWithRetry(ctx context.Context, txID string, count int) (string, error) {
	if count >= 10 {
		err := errors.New("max retry exceeded")
		log.Error("max retry exceeded", err)
		return "", err
	}
	res, err := c.svc.GetData(ctx, txID)
	if err != nil {
		log.Error("transaction get failed", err)
		return "", err
	}
	if res == "" {
		return c.downloadWithRetry(ctx, txID, count+1)
	}
	return res, nil
}

// UploadAdvertisementMetadata upload metadata
func (c *Client) UploadAdvertisementMetadata(ctx context.Context, advertisement ad.Advertisement) (string, error) {
	return c.uploadJSON(ctx, advertisement)
}

// UploadPostMetadata upload metadata
func (c *Client) UploadPostMetadata(ctx context.Context, post post.Post) (string, error) {
	return c.uploadJSON(ctx, post)
}

func (c *Client) uploadJSON(ctx context.Context, js interface{}) (string, error) {
	b, err := json.Marshal(&js)
	if err != nil {
		log.Error("json marshal failed", err)
		return "", err
	}
	return c.sendTransaction(ctx, b, "application/json")
}

// ToURL tx to url
func (c *Client) ToURL(tx string) string {
	return toArweaveURL(tx)
}

func toArweaveURL(imageTx string) string {
	return uri + "/" + imageTx
}

func urlToTxID(url string) string {
	return strings.ReplaceAll(url, uri+"/", "")
}

func (c *Client) sendTransaction(ctx context.Context, data []byte, contenttype string) (string, error) {
	ar, err := newTransact()
	if err != nil {
		log.Error("send transaction failed", err)
		return "", err
	}
	txBuilder, err := ar.CreateTransaction(ctx, c.w, "0", data, "")
	if err != nil {
		log.Error("create transaction failed", err)
		return "", err
	}
	txBuilder.AddTag("Content-Type", contenttype)
	txn, err := txBuilder.Sign(c.w)
	if err != nil {
		log.Error("sign transaction failed", err)
		return "", err
	}
	_, err = ar.SendTransaction(ctx, txn)
	if err != nil {
		log.Error("send transaction failed", err)
		return "", err
	}
	return txBuilder.Hash(), nil
}
