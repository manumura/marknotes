// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package pub_elements

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "html/template"

func PostPreview(content template.HTML) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"text-start rounded-xl lg:max-w-3xl md:max-w-md md:basis-7/12 md:mr-8 md:ml-0 md:p-3 md:my-8 xs:p-3 xs:m-4 bg-stone-100 px-6 dark:bg-slate-950 \"><article class=\"prose prose-slate lg:prose-xl md:prose-lg dark:prose-invert prose-h2:bg-clip-text prose-h2:text-transparent prose-h2:bg-gradient-to-r prose-h2:from-pink-600 prose-h2:via-purple-600 prose-h2:to-indigo-600 prose-h3:bg-clip-text  prose-h3:text-transparent  prose-h3:bg-gradient-to-r  prose-h3:from-teal-600  prose-h3:via-sky-600  prose-h3:to-indigo-600\" id=\"preview\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = RenderRaw(content).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</article></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func RenderRaw(content template.HTML) templ.Component {
	return templ.Raw[template.HTML](content)
}
