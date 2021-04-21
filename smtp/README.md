golang邮件发送模块，带附件

config := `{"username":"testx@gmail.com","password":"testx","host":"smtp.gmail.com","port":587}`
mail := NewEMail(config)
if mail.Username != "testx@gmail.com" {
    t.Fatal("email parse get username error")
}
if mail.Password != "testx" {
    t.Fatal("email parse get password error")
}
if mail.Host != "smtp.gmail.com" {
    t.Fatal("email parse get host error")
}
if mail.Port != 587 {
    t.Fatal("email parse get port error")
}
mail.To = []string{"abc@gmail.com"}
mail.From = "testx@gmail.com"
mail.Subject = "hi, just from beego!"
mail.Text = "Text Body is, of course, supported!"
mail.HTML = "<h1>Fancy Html is supported, too!</h1>"
mail.AttachFile("/Users/testx/github/beego/beego.go")
mail.Send()