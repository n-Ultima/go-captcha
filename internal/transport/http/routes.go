package http

import (
	"encoding/json"
	"image"
	"image/draw"
	"image/png"
	"net/http"
	"os"
	"path/filepath"

	"github.com/PolyLmao/go-captcha/internal/models"
	"github.com/goki/freetype/truetype"
	"github.com/gorilla/mux"
	"github.com/leonelquinteros/gorand"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// swagger:parameters serveImage
type captchaCodeParameter struct {
	// The captcha code that will lead us to an image.
	// in: path
	// required: true
	Code string
}

// swagger:route GET /images/{code} Captcha ServeImage
// Serve a captcha image by captcha code.
// responses:
//  200: captchaImage
func (h *Handler) ServeImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "image/png")
	name := mux.Vars(r)["code"]
	http.ServeFile(w, r, "images/"+name+".png")
}

// Serve our swagger specification for use by our documentation.
func (h *Handler) ServeDocs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "swagger.yml")
}

// swagger:route GET /captcha Captcha NewCaptcha
// Create a new captcha code and image, then return a response object containing our newly created captcha.
// responses:
//  200: NewCaptcha
func (h *Handler) NewCaptcha(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	randString, _ := createImage()
	json.NewEncoder(w).Encode(models.Captcha{Image: r.Host + "/images/" + randString, Code: randString})
}

// Create a new image with a randomized captcha code, and then return the code.
func createImage() (string, error) {
	fontFile, err := os.ReadFile("assets/font/font.ttf")
	if err != nil {
		return "", err
	}
	capFont, err := truetype.Parse(fontFile)
	if err != nil {
		return "", err
	}
	dst := image.NewRGBA(image.Rect(0, 0, 200, 50))
	draw.Draw(dst, dst.Bounds(), image.White, image.Point{}, draw.Src)
	drawer := &font.Drawer{
		Dst: dst,
		Src: image.Black,
		Face: truetype.NewFace(capFont, &truetype.Options{
			Size: 24,
			DPI:  72,
		}),
		Dot: fixed.P(20, 30),
	}
	randString, _ := gorand.GetAlphaNumString(8)
	drawer.DrawString(randString)
	outFile, err := os.Create(filepath.Join("images", randString+".png"))
	if err != nil {
		return "", err
	}
	defer outFile.Close()
	if err := png.Encode(outFile, dst); err != nil {
		return "", err
	}
	return randString, nil
}
