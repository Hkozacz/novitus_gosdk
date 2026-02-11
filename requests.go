package novitus_gosdk

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Document interface {
	Validate() error
}

type DiscountMarkup struct {
	Type  string `json:"type"` //Enum: "percent_discount" "percent_markup" "value_discount" "value_markup"
	Name  string `json:"name,omitempty"`
	Value string `json:"value"`
}

type Summary struct {
	DiscountMarkup *DiscountMarkup `json:"discount_markup,omitempty"`
	Total          string          `json:"total,omitempty"`
	PayIn          string          `json:"pay_in,omitempty"`
	Change         string          `json:"change,omitempty"`
}

type EDocument struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
	PrintSendMode string `json:"print_send_mode,omitempty"`
}

type Buyer struct {
	Name      string   `json:"name,omitempty"`
	IdType    string   `json:"id_type,omitempty"`
	Id        string   `json:"id,omitempty"`
	LabelType string   `json:"label_type,omitempty"`
	Address   []string `json:"address,omitempty"`
	Nip       string   `json:"nip,omitempty"`
	EDocument `json:"e_document,omitempty"`
}

type SystemInfo struct {
	CashierName  string `json:"cashier_name,omitempty"`
	CashNumer    string `json:"cash_number,omitempty"`
	SystemNumber string `json:"system_number,omitempty"`
}

type DeviceControl struct {
	OpenDrawer        bool   `json:"open_drawer,omitempty"`
	FeedAfterPrintout bool   `json:"feed_after_printout,omitempty"`
	PaperCut          string `json:"paper_cut,omitempty"`
}

type Receipt struct {
	Items         []interface{}              `json:"items,omitempty"` // Required: true
	Payments      []interface{}              `json:"payments,omitempty"`
	Summary       `json:"summary,omitempty"` // Required: true
	PrintoutLines []interface{}              `json:"printout_lines,omitempty"`
	Buyer         *Buyer                     `json:"buyer,omitempty"`
	SystemInfo    *SystemInfo                `json:"system_info,omitempty"`
	DeviceControl *DeviceControl             `json:"device_control,omitempty"`
}

func (r *Receipt) Validate() error {
	if len(r.Items) == 0 {
		return fmt.Errorf("items are required")
	}
	if r.Summary.Total == "" {
		return fmt.Errorf("summary.total is required")
	}
	return nil
}

type Info struct {
	Number        string `json:"number,omitempty"`
	CopyCount     int    `json:"copy_count,omitempty"`
	DateOfSell    string `json:"date_of_sell,omitempty"`
	DateOfPayment string `json:"date_of_payment,omitempty"`
	PaymentForm   string `json:"payment_form,omitempty"`
	Paid          string `json:"paid,omitempty"`
}

type TransactionSide struct {
	Name      string `json:"name,omitempty"`
	PrintInfo string `json:"print_info,omitempty"` // Enum: "place_for_signature" "name_and_place_for_signature" "none"
}

func (ts *TransactionSide) Validate() error {
	if ts.PrintInfo != "" && ts.PrintInfo != "place_for_signature" && ts.PrintInfo != "name_and_place_for_signature" && ts.PrintInfo != "none" {
		return fmt.Errorf("print_info must be one of: place_for_signature, name_and_place_for_signature, none")
	}
	return nil
}

type Options struct {
	SkipDescriptionValueToPay                  bool `json:"skip_description_value_to_pay"`
	SkipBlockGrossValueInAccountingTax         bool `json:"skip_block_gross_value_in_accounting_tax"`
	BuyerBold                                  bool `json:"buyer_bold"`
	SellerBold                                 bool `json:"seller_bold"`
	BuyerNipBold                               bool `json:"buyer_nip_bold"`
	SellerNipBold                              bool `json:"seller_nip_bold"`
	PrintLabelDescriptionSymbolInInvoiceHeader bool `json:"print_label_description_symbol_in_invoice_header"`
	PrintPositionNumberInInvoiceHeader         bool `json:"print_position_number_in_invoice_header"`
	PrintPositionNumberInvoice                 bool `json:"print_position_number_invoice"`
	ToPayLabelBeforeAcountingTaxBlock          bool `json:"to_pay_label_before_acounting_tax_block"`
	PrintCentsInWords                          bool `json:"print_cents_in_words"`
	DontPrintSellDateIfEqualCreateDate         bool `json:"dont_print_sell_date_if_equal_create_date"`
	DontPrintSellerDataInHeader                bool `json:"dont_print_seller_data_in_header"`
	DontPrintSellItemsDescription              bool `json:"dont_print_sell_items_description"`
	EnablePaymentForm                          bool `json:"enable_payment_form"`
	DontPrintCustomerData                      bool `json:"dont_print_customer_data"`
	PrintPaydInCash                            bool `json:"print_payd_in_cash"`
	SkipSellerLabel                            bool `json:"skip_seller_label"`
	PrintInvoiceTaxLabel                       bool `json:"print_invoice_tax_label"`
}

type AdditionalInfo struct {
	Text          string `json:"text,omitempty"`
	Bold          bool   `json:"bold,omitempty"`
	Justification string `json:"justification,omitempty"` // Enum: "left" "center" "right"
}

type Invoice struct {
	Info           `json:"info"`    // Required: true
	Buyer          `json:"buyer"`   // Required: true
	Recipient      *TransactionSide `json:"recipient,omitempty"`
	Seller         *TransactionSide `json:"seller,omitempty"`
	Options        `json:"options"`
	Items          []interface{}    `json:"items"` // Required: true
	Payments       []interface{}    `json:"payments,omitempty"`
	Summary        `json:"summary"` // Required: true
	PrintoutLines  []interface{}    `json:"printout_lines,omitempty"`
	AdditionalInfo []AdditionalInfo `json:"additional_info,omitempty"`
	DeviceControl  `json:"device_control,omitempty"`
	SystemInfo     `json:"system_info,omitempty"`
}

func (i *Invoice) Validate() error {
	if i.Info.Number == "" {
		return fmt.Errorf("info.number is required")
	}
	if len(i.Items) == 0 {
		return fmt.Errorf("items are required")
	}
	if i.Summary.Total == "" {
		return fmt.Errorf("summary.total is required")
	}
	if i.Buyer.Name == "" && i.Buyer.Nip == "" {
		return fmt.Errorf("buyer.name or buyer.nip is required")
	}
	return nil
}

type PrintoutOptions struct {
	WithoutHeader    bool `json:"without_header,omitempty"`
	LeftMargin       bool `json:"left_margin,omitempty"`
	CopyOnly         bool `json:"copy_only,omitempty"`
	FiscalMarginsOff bool `json:"fiscal_margins_off,omitempty"`
}

type Printout struct {
	Options        *PrintoutOptions `json:"options,omitempty"`
	Lines          []interface{}    `json:"lines"` // Required: true
	*EDocument     `json:"e_document,omitempty"`
	*SystemInfo    `json:"system_info,omitempty"`
	*DeviceControl `json:"device_control,omitempty"`
}

func (p *Printout) Validate() error {
	if len(p.Lines) == 0 {
		return fmt.Errorf("lines are required")
	}
	return nil
}

// Items

type Article struct {
	Name           string          `json:"name"`                      // Required: true
	PTU            string          `json:"ptu"`                       // Enum: "A" - "G" Required: true https://www.posnet.com.pl/gdzie-kupic
	Quantity       string          `json:"quantity"`                  // Quantity in units, e.g. "1.00" Required: true
	Price          string          `json:"price"`                     // Price in currency, e.g. "1.00" Required: true
	Value          string          `json:"value"`                     // Total value for the item, e.g. "1.00" Required: true
	Unit           string          `json:"unit,omitempty"`            // Enum: "szt" - "kg", etc.
	DiscountMarkup *DiscountMarkup `json:"discount_markup,omitempty"` // Optional, e.g. "0.00"
	Code           string          `json:"code,omitempty"`            // Optional, e.g. "1234567890123" Can be set only if Description is not Set
	Description    string          `json:"description,omitempty"`     // Optional, e.g. "Sample Item"
}

func (a *Article) Validate() error {
	if a.Name == "" {
		return fmt.Errorf("name is required")
	}
	if a.PTU != "A" && a.PTU != "B" && a.PTU != "C" && a.PTU != "D" && a.PTU != "E" && a.PTU != "F" && a.PTU != "G" {
		return fmt.Errorf("ptu must be one of: A, B, C, D, E, F, G")
	}
	if a.Quantity == "" {
		return fmt.Errorf("quantity is required")
	}
	if a.Price == "" {
		return fmt.Errorf("price is required")
	}
	if a.Value == "" {
		return fmt.Errorf("value is required")
	}
	if a.Unit != "" && a.Unit != "szt" && a.Unit != "kg" {
		return fmt.Errorf("unit must be one of: szt, kg, etc.")
	}
	decValue, err := decimal.NewFromString(a.Value)
	if err != nil {
		return fmt.Errorf("invalid value format: %w", err)
	}
	decPrice, err := decimal.NewFromString(a.Price)
	if err != nil {
		return fmt.Errorf("invalid price format: %w", err)
	}
	decQuantity, err := decimal.NewFromString(a.Quantity)
	if err != nil {
		return fmt.Errorf("invalid quantity format: %w", err)
	}
	if decValue != decPrice.Mul(decQuantity) {
		return fmt.Errorf("value must be equal to price multiplied by quantity")
	}
	return nil
}

type Advance struct {
	Description string `json:"description,omitempty"` // Required: true, e.g. "Advance Payment"
	PTU         string `json:"ptu"`                   // Enum: "A" - "G" Required: true https://www.posnet.com.pl/gdzie-kupic
	Value       string `json:"value"`                 // Value in currency, e.g. "100.00" Required: true
}

func (a *Advance) Validate() error {
	if a.Description == "" {
		return fmt.Errorf("description is required")
	}
	if a.PTU != "A" && a.PTU != "B" && a.PTU != "C" && a.PTU != "D" && a.PTU != "E" && a.PTU != "F" && a.PTU != "G" {
		return fmt.Errorf("ptu must be one of: A, B, C, D, E, F, G")
	}
	if a.Value == "" {
		return fmt.Errorf("value is required")
	}
	return nil
}

type AdvanceReturn struct {
	Description string `json:"description,omitempty"` // Required: true, e.g. "Advance Return"
	PTU         string `json:"ptu"`                   // Enum: "A" - "G" Required: true https://www.posnet.com.pl/gdzie-kupic
	Value       string `json:"value"`                 // Value in currency, e.g. "50.00" Required: true
}

func (a *AdvanceReturn) Validate() error {
	if a.Description == "" {
		return fmt.Errorf("description is required")
	}
	if a.PTU != "A" && a.PTU != "B" && a.PTU != "C" && a.PTU != "D" && a.PTU != "E" && a.PTU != "F" && a.PTU != "G" {
		return fmt.Errorf("ptu must be one of: A, B, C, D, E, F, G")
	}
	if a.Value == "" {
		return fmt.Errorf("value is required")
	}
	return nil
}

type Container struct {
	Name     string `json:"name,omitempty"`     // e.g. "Container Name"
	Number   string `json:"number,omitempty"`   // e.g. "12345"
	Quantity string `json:"quantity,omitempty"` // Quantity in units, e.g. "10.00"
	Value    string `json:"value"`              // Total value for the container, e.g. "100.00" Required: true
}

func (c *Container) Validate() error {
	if c.Value == "" {
		return fmt.Errorf("value is required")
	}
	return nil
}

type ContainerReturn struct {
	Name     string `json:"name"`     // e.g. "Container Name"
	Number   string `json:"number"`   // e.g. "12345"
	Quantity string `json:"quantity"` // Quantity in units, e.g. "10.00"
	Value    string `json:"value"`    // Total value for the container, e.g. "100.00" Required: true
}

func (cr *ContainerReturn) Validate() error {
	if cr.Value == "" {
		return fmt.Errorf("value is required")
	}
	return nil
}

// Payments

type Cash struct {
	Value string `json:"value"` // Value in currency, e.g. "100.00" Required: true
}

func (c *Cash) Validate() error {
	if c.Value == "" {
		return fmt.Errorf("value is required")
	}
	return nil
}

type TypicalPaymentMethod struct {
	Name  string `json:"name,omitempty"` // enum "card", cheque, coupon, other, credit, account, transfer, mobile, voucher
	Value string `json:"value"`          // Value in currency, e.g. "100.00" Required: true
}

func (t *TypicalPaymentMethod) Validate() error {
	if t.Value == "" {
		return fmt.Errorf("value is required")
	}
	if t.Name != "card" && t.Name != "cheque" && t.Name != "coupon" && t.Name != "other" && t.Name != "credit" && t.Name != "account" && t.Name != "transfer" && t.Name != "mobile" && t.Name != "voucher" {
		return fmt.Errorf("name must be one of: card, cheque, coupon, other, credit, account, transfer, mobile, voucher")
	}
	return nil
}

type Currency struct {
	Course        string `json:"course"`         // e.g. "1.00" Required: true
	CurrencyValue string `json:"currency_value"` // e.g. "USD" Required: true
	LocalValue    string `json:"local_value"`    // e.g. "100.00" Required: true
	IsChange      bool   `json:"is_change"`      // true if this is a change, false otherwise Required: true
	Name          string `json:"name"`           // e.g. "USD" Required: true
}

func (c *Currency) Validate() error {
	if c.Course == "" {
		return fmt.Errorf("course is required")
	}
	if c.CurrencyValue == "" {
		return fmt.Errorf("currency_value is required")
	}
	if c.LocalValue == "" {
		return fmt.Errorf("local_value is required")
	}
	if c.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}

// Printout Lines

type PrintoutLine struct {
	Text   string `json:"text"`   // The text to be printed, e.g. "Sample Text" Required: true
	Masked bool   `json:"masked"` // true if the text should be masked, false otherwise Required: true
}

func (p *PrintoutLine) Validate() error {
	if p.Text == "" {
		return fmt.Errorf("text is required")
	}
	return nil
}

type TextLine struct {
	Bold       bool   `json:"bold,omitempty"`        // true if the text should be bold, false otherwise
	Invers     bool   `json:"invers,omitempty"`      // true if the text should be inverted, false otherwise
	Center     bool   `json:"center,omitempty"`      // true if the text should be centered, false otherwise
	FontNumber int    `json:"font_number,omitempty"` // Font number, e.g. 1, 2, 3
	Big        bool   `json:"big,omitempty"`         // true if the text should be big, false otherwise
	Height     int    `json:"height,omitempty"`      // Height of the text in points, e.g. 12
	Width      int    `json:"width,omitempty"`       // Width of the text in points, e.g. 100
	Text       string `json:"text"`                  // The text to be printed, e.g. "Sample Text" Required: true
	Masked     bool   `json:"masked"`                // true if the text should be masked, false otherwise Required: true
}

func (tl *TextLine) Validate() error {
	if tl.Text == "" {
		return fmt.Errorf("text is required")
	}
	if tl.Height < 0 {
		return fmt.Errorf("height must be a positive integer")
	}
	if tl.Width < 0 {
		return fmt.Errorf("width must be a positive integer")
	}
	if tl.FontNumber < 1 || tl.FontNumber > 3 {
		return fmt.Errorf("font_number must be between 1 and 3")
	}
	return nil
}
