// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Filters(subcategories []string) templ.Component {
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
		templ_7745c5c3_Err = filterForm(subcategories).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func filterForm(subcategories []string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p>Filter your search results</p><main id=\"filter-form\"><section id=\"variants\"><label for=\"variants-input\">Variants</label> <input type=\"text\" id=\"variants-input\" name=\"variants\" placeholder=\"Ex: 1oz,1g\"></section><section id=\"filter-subcategory\"><label for=\"subcategory-input\">Subcategory</label> <select name=\"sub\" id=\"subcategory-input\"><option value=\"\">Select Subcategory</option> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, sub := range subcategories {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(sub))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(sub)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/filters.templ`, Line: 23, Col: 47}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></section><section id=\"minPricePerG\"><label for=\"minPricePerG-input\">Min Price Per G</label> <input type=\"text\" id=\"minPricePerG-input\" name=\"min_price_per_g\" placeholder=\"Ex: 0\"></section><section id=\"maxPricePerG\"><label for=\"maxPricePerG-input\">Max Price Per G</label> <input type=\"text\" id=\"maxPricePerG-input\" name=\"max_price_per_g\" placeholder=\"Ex: 100\"></section><section id=\"minPrice\"><label for=\"minPrice-input\">Min Price</label> <input type=\"text\" id=\"minPrice-input\" name=\"min_price\" placeholder=\"Ex: 0\"></section><section id=\"maxPrice\"><label for=\"maxPrice-input\">Max Price</label> <input type=\"text\" id=\"maxPrice-input\" name=\"max_price\" placeholder=\"Ex: 100\"></section><section id=\"excluded-brands\"><label for=\"excluded-brands-input\">Excl. Brands</label> <input type=\"text\" id=\"excluded-brands-input\" name=\"not_brands\" placeholder=\"Ex: select,curaleaf\"></section><section id=\"included-brands\"><label for=\"included-brands-input\">Incl. Brands</label> <input type=\"text\" id=\"included-brands-input\" name=\"brands\" placeholder=\"Ex: grassroots,cresco\"></section><section id=\"excluded-terms\"><label for=\"excluded-terms-input\">Excl. Terms</label> <input type=\"text\" id=\"excluded-terms-input\" name=\"exclude\" placeholder=\"Ex: distillate,C02\"></section><section id=\"included-terms\"><label for=\"included-terms-input\">Incl. Terms</label> <input type=\"text\" id=\"included-terms-input\" name=\"include\" placeholder=\"Ex: resin,live\"></section></main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
