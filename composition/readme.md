# Golang composition example

Here we got an example that in golang there is no inheritance. Only a composition.
Enchanted sword compose 'Sword' and all its methods.
But when we invoke EnchantedSword.String method, it actually call Sword.String, which calls Sword.Damage();

# go-test-report

Also there is an example of go-test-report usage
```shell
go test -json | go-test-report
```