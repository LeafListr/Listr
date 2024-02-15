// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/Linkinlog/LeafListr/internal/api/models"

func Index(title string, dispensaries []string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><link rel=\"icon\" type=\"image/svg+xml\" href=\"/assets/logo.png\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta name=\"theme-color\" media=\"(prefers-color-scheme: dark)\" content=\"black\"><script src=\"https://unpkg.com/htmx.org@1.9.10\" integrity=\"sha384-D1Kt99CQMDuVetoL1lrYwg5t+9QdHe7NLX/SoJYkXDFfX37iInKRy5xLSi8nO7UC\" crossorigin=\"anonymous\"></script><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/index.templ`, Line: 13, Col: 17}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title></head><body><main hx-include=\"[name=&#39;category&#39;], [name=&#39;dispensary&#39;], [name=&#39;location&#39;], [name=&#39;recreational&#39;]\" hx-params=\"*\" id=\"app\"><h2 id=\"now-viewing\">Please select a dispensary</h2><div class=\"toggle-switch\"><input hx-get=\"/views/nav\" hx-target=\"#nav\" type=\"checkbox\" id=\"med-rec-toggle\" name=\"recreational\"> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 = []any{"switch-label"}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var3...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label for=\"med-rec-toggle\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var3).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><span class=\"toggle-text\">Recreational</span> <span class=\"switch-handle\"></span> <span class=\"toggle-text\">Medical</span></label></div><nav id=\"nav\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Dispensaries(dispensaries).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<br>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Locations("", []*models.Location{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</nav><h3 id=\"loading-categories\" class=\"htmx-indicator\">Loading Categories...<img src=\"/assets/blocks-wave.svg\"></h3><section id=\"offers-and-search\" hx-target=\"#products\" hx-indicator=\"#loading-products\" hx-swap=\"outerHTML\"></section><h3 id=\"loading-products\" class=\"htmx-indicator\">Loading Products...<img src=\"/assets/blocks-wave.svg\"></h3>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 = []any{"products-grid"}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var4...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"products\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var4).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></main></body></html><style>\n        :root {\n          font-family: Inter, system-ui, Avenir, Helvetica, Arial, sans-serif;\n          line-height: 1.5;\n          font-weight: 400;\n\n          color-scheme: dark;\n          color: rgba(255, 255, 255, 0.87);\n          background-color: #242424;\n\n          font-synthesis: none;\n          text-rendering: optimizeLegibility;\n          -webkit-font-smoothing: antialiased;\n          -moz-osx-font-smoothing: grayscale;\n        }\n\n        a {\n          font-weight: 500;\n          color: #646cff;\n          text-decoration: inherit;\n        }\n\n        a:hover {\n          color: #535bf2;\n        }\n\n        body {\n          margin: 0;\n          display: flex;\n          place-items: center;\n          min-width: 320px;\n          min-height: 100vh;\n        }\n\n        h1 {\n          font-size: 3.2em;\n          line-height: 1.1;\n        }\n\n        .card {\n          padding: 2em;\n        }\n\n        #app {\n          max-width: 1280px;\n          margin: 0 auto;\n          text-align: center;\n        }\n\n        #categories label {\n          display: inline-block;\n          border-radius: 8px;\n          border: 1px solid transparent;\n          padding: 0.6em 1.2em;\n          font-size: 1em;\n          font-weight: 500;\n          font-family: inherit;\n          background-color: #1a1a1a;\n          cursor: pointer;\n          transition: border-color 0.25s;\n        }\n\n        #categories label:hover {\n          border-color: #646cff;\n        }\n        #categories label:focus,\n        #categories label:focus-visible {\n          outline: 4px auto -webkit-focus-ring-color;\n        }\n\n        button {\n          grid-area: submit;\n          border-radius: 8px;\n          border: 1px solid transparent;\n          padding: 0.6em 1.2em;\n          font-size: 1em;\n          font-weight: 500;\n          font-family: inherit;\n          background-color: #1a1a1a;\n          cursor: pointer;\n          transition: border-color 0.25s;\n        }\n        button:hover {\n          border-color: #646cff;\n        }\n        button:focus,\n        button:focus-visible {\n          outline: 4px auto -webkit-focus-ring-color;\n        }\n\n        nav select {\n          min-width: 275px;\n        }\n        .products-grid {\n          display: grid;\n          gap: 16px;\n          grid-template-columns: repeat(auto-fill, minmax(275px, 1fr));\n        }\n\n        @media (min-width: 992px) {\n          .products-grid {\n            grid-template-columns: repeat(3, minmax(275px, 1fr));\n          }\n        }\n\n        @media (max-width: 600px) {\n          .products-grid {\n            grid-template-columns: repeat(2, minmax(275px, 1fr));\n          }\n        }\n\n        @media (max-width: 480px) {\n          .products-grid {\n            grid-template-columns: repeat(1, minmax(275px, 1fr));\n          }\n        }\n        .product-card {\n          display: grid;\n          grid-template-columns: 1fr 1fr;\n          grid-template-rows: auto auto auto;\n          grid-auto-rows: minmax(min-content, max-content);\n          grid-template-areas:\n            \"headline headline headline\"\n            \"subcategory gram-price variant\"\n            \"cannabinoids cannabinoids cannabinoids\"\n            \"terpenes terpenes terpenes\";\n          border: 1px solid #ccc;\n          border-radius: 8px;\n          box-shadow: rgba(0, 0, 0, 0.56) 0px 22px 70px 4px;\n          background-color: #2c2030;\n          padding: 20px;\n          text-align: center;\n          align-items: center;\n        }\n\n        .product-card p {\n            font-size: .75rem;\n        }\n\n        .product-card img {\n          grid-area: image;\n          width: 64px;\n          height: 64px;\n          border-radius: 4px;\n        }\n\n        .product-card .product-headline {\n          grid-area: headline;\n          grid-template-areas: \"image name\";\n          display: grid;\n        }\n\n        .product-card .product-name {\n          grid-area: name;\n          margin: 0 0 0 10px;\n        }\n\n        .product-card .product-cannabinoids {\n          grid-area: cannabinoids;\n          display: grid;\n          grid-gap: 10px;\n          margin: 6px 4px 0 0;\n        }\n\n        .product-card .product-cannabinoids p {\n            font-size: .55rem;\n        }\n\n        .product-card .product-terpenes {\n          grid-area: terpenes;\n          display: grid;\n          grid-gap: 10px;\n          grid-template-columns: repeat(2, 1fr);\n          margin: 6px 4px 0 0;\n        }\n\n        .product-card .product-terpenes p {\n            font-size: .55rem;\n        }\n\n        .product-card .product-subcategory {\n          grid-area: subcategory;\n        }\n\n        .product-card .product-gram-price {\n          grid-area: gram-price;\n        }\n\n        .product-card .product-variants {\n          grid-area: variant;\n          padding: 12px 0;\n        }\n\n        label {\n          font-size: 12px;\n          font-weight: bold;\n        }\n\n        .discounted-price {\n          color: green;\n        }\n\n        .discounted {\n          color: red;\n          text-decoration: line-through;\n          font-size: 0.65rem;\n          line-height: 0;\n          vertical-align: super;\n        }\n        #categories {\n          display: grid;\n          grid-template-columns: repeat(2, 1fr);\n          gap: 10px;\n          margin: 10px;\n        }\n        @media (min-width: 768px) {\n          #categories {\n            grid-template-columns: repeat(4, 1fr);\n          }\n        }\n        #categories button {\n          padding: 10px;\n          margin: 0 2px;\n          border: 1px solid #858585;\n          border-radius: 5px;\n          cursor: pointer;\n        }\n        #categories input[type=\"radio\"] {\n          position: fixed;\n          opacity: 0;\n          pointer-events: none;\n        }\n\n        .accordion {\n          margin: 1rem 0;\n        }\n\n        .header {\n          display: flex;\n          width: 100%;\n        }\n\n        .header .nav {\n          flex: 1;\n          margin-bottom: 7px;\n        }\n\n        .details {\n          box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;\n          background-color: #1d1d20;\n          border-radius: 60px;\n          padding: 1rem;\n          opacity: 0;\n          visibility: hidden;\n          transition:\n            opacity 0.5s,\n            visibility 0s 0.5s,\n            max-height 0.5s ease;\n          max-height: 0;\n          overflow: hidden;\n        }\n\n        .details.open {\n          opacity: 1;\n          visibility: visible;\n          max-height: 100%;\n          transition:\n            opacity 0.5s,\n            visibility 0s,\n            max-height 0.5s ease;\n        }\n\n        #filter-form {\n          display: grid;\n          grid-template-areas:\n            \"variants subcategory\"\n            \"minPrice maxPrice\"\n            \"excluded-brands included-brands\"\n            \"excluded-terms included-terms\"\n            \"submit submit\";\n          gap: 10px;\n          margin: 7px;\n          padding-bottom: 7px;\n        }\n        #filter-form input {\n          width: 90%;\n        }\n        #variants {\n          grid-area: variants;\n        }\n        #filter-form filter-subcategory {\n          grid-area: subcategory;\n        }\n        #minPrice {\n          grid-area: minPrice;\n        }\n        #maxPrice {\n          grid-area: maxPrice;\n        }\n        .toggle-switch {\n          display: inline-block;\n          position: relative;\n        }\n\n        #med-rec-toggle {\n          display: none;\n        }\n\n        .switch-label {\n          display: block;\n          cursor: pointer;\n          background-color: #ddd;\n          border-radius: 20px;\n          overflow: hidden;\n          text-align: center;\n        }\n\n        .switch-handle {\n          display: block;\n          background-color: #242424;\n          position: absolute;\n          border-radius: 20px;\n          top: 2px;\n          left: 2px;\n          right: 50%;\n          bottom: 2px;\n          transition: all 0.3s ease;\n        }\n\n        #med-rec-toggle:checked + .switch-label .switch-handle {\n          right: 2px;\n          left: 50%;\n        }\n\n        .toggle-text {\n          line-height: 30px;\n          padding: 0 20px;\n          display: inline-block;\n          width: 100px;\n          color: #000;\n          font-weight: bold;\n        }\n\n        .switch-label {\n          box-shadow: 0 0 5px rgba(0, 0, 0, 0.3);\n        }\n\n        #sorter-form {\n          display: grid;\n          grid-template-areas:\n            \"gram-price-sort-method terps\"\n            \"price-sort-method terps\"\n            \"thc-sort-method terps\"\n            \"terpene-sort-method terps\"\n            \"submit submit\";\n          grid-template-columns: 1fr 1fr;\n          gap: 10px;\n          margin: 7px;\n          padding-bottom: 7px;\n        }\n        #sorter-form select {\n          width: 100%;\n        }\n\n        #gram-price-sort-method {\n          grid-area: gram-price-sort-method;\n        }\n\n        #price-sort-method {\n          grid-area: price-sort-method;\n        }\n\n        #terps {\n          grid-area: terps;\n        }\n        .htmx-indicator {\n          display: none;\n        }\n        .htmx-request .htmx-indicator {\n          display: block;\n        }\n        .htmx-request.htmx-indicator {\n          display: block;\n        }\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
