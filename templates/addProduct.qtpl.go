// Code generated by qtc from "addProduct.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// Page of product add
//

//line ../templates/addProduct.qtpl:3
package templates

//line ../templates/addProduct.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line ../templates/addProduct.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line ../templates/addProduct.qtpl:4
type ProductAddPage struct {
}

//line ../templates/addProduct.qtpl:8
func (p *ProductAddPage) StreamBody(qw422016 *qt422016.Writer) {
//line ../templates/addProduct.qtpl:8
	qw422016.N().S(`
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
</head>

<body>
    <h5 id="error"></h5>

    <form id="form">
        <input type="text" name="name" placeholder="Name">
        <input type="number" name="price" placeholder="Price">

        <button type="submit">Создать</button>
    </form>

    <a href="/products">
        <button type="button">Отмена</button>
    </a>
</body>

<script>
    var form = document.getElementById('form');
    form.addEventListener('submit', function(event) {
        event.preventDefault();

        var formData = new FormData(form), result = {};
        for (var entry of formData.entries()) {
            result[entry[0]] = entry[1];
        }

        var xmlHttp = new XMLHttpRequest();
        xmlHttp.open("POST", "/cmd/add-product", false);
        xmlHttp.send(JSON.stringify(result));

        var jsonData = JSON.parse(xmlHttp.responseText);
        if (jsonData.status === "failure") {
            document.getElementById("error").innerHTML = jsonData.error;
        } else {
            window.location.replace("/products");
        }
    });
</script>

</html>
`)
//line ../templates/addProduct.qtpl:54
}

//line ../templates/addProduct.qtpl:54
func (p *ProductAddPage) WriteBody(qq422016 qtio422016.Writer) {
//line ../templates/addProduct.qtpl:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line ../templates/addProduct.qtpl:54
	p.StreamBody(qw422016)
//line ../templates/addProduct.qtpl:54
	qt422016.ReleaseWriter(qw422016)
//line ../templates/addProduct.qtpl:54
}

//line ../templates/addProduct.qtpl:54
func (p *ProductAddPage) Body() string {
//line ../templates/addProduct.qtpl:54
	qb422016 := qt422016.AcquireByteBuffer()
//line ../templates/addProduct.qtpl:54
	p.WriteBody(qb422016)
//line ../templates/addProduct.qtpl:54
	qs422016 := string(qb422016.B)
//line ../templates/addProduct.qtpl:54
	qt422016.ReleaseByteBuffer(qb422016)
//line ../templates/addProduct.qtpl:54
	return qs422016
//line ../templates/addProduct.qtpl:54
}