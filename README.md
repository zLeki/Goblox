<div id="top"></div>

<!-- PROJECT SHIELDS -->


<!-- ![Visitors](https://estruyf-github.azurewebsites.net/api/VisitorHit?user=wst24365888&repo=ez4o/convert-json-cli&countColor=rgb(0,%20126,%20198)) -->

<br />

![Goblox](https://i.ibb.co/WfCGTbX/Untitled.png)

<!-- PROJECT LOGO -->
<br />
<div align="center">
<p align="center">
    <a href="https://github.com/zLeki/Goblox#usage"><strong>Explore Usage »</strong></a>
    <br />
    <br />
    <a href="https://github.com/zLeki/Goblox/issues">Report Bug</a>
    ·
    <a href="https://github.com/zLeki/Goblox/issues">Request Feature</a>
  </p>
</div>



<!-- ABOUT THE PROJECT -->

## About The Project


**Goblox** is a great API Wrapper for roblox api made by developers for developers



<!-- GETTING STARTED -->

## Getting Started

### Download
# Requires Go 1.18 (not sure if it works on lower versions)

`go get github.com/zLeki/Goblox`
### Example

```go
package main
import (
	"github.com/zLeki/Goblox/account"
	"github.com/zLeki/Goblox/groups"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
)
func snipe(target int) {
	f, _ := ioutil.ReadFile("ids.txt")
	list := strings.Split(string(f), "\n")
	for _, v := range list {
		id, _ := strconv.Atoi(v)
		price := catalog.GetInfo(id, acc).Data[0].LowestPrice
		log.Println(price)
		if price <= target {
			log.Println("Cheaper then target should purchase Price: ", price)
		}else{
			log.Println("Higher then target should not buy. Price: ", price)
		}
	}}
func main() {
	acc := account.Validate("cookie")
	log.Println(groups.GetUserRoleInGroup(439403718, 13644831, acc))
	_, err := groups.SendShout(acc, "13644831", "Hello, World!")
	if err != nil {
		log.Println(err)
	}
	err2 := groups.UserPayout("439403718", 13644831, 5, acc)
	if err2 != nil {
		log.Println(err2)
	}
	gi := groups.Groupinfo("13644831")
	log.Println(gi.Shout.Content, gi.MemberCount, gi.Description)
	for _, v := range groups.GetRoles(13644831) {
		log.Println(v.Name)
	}
    log.Println(profile.Info("xLeki").DisplayName)
    snipe(5)
}
```


## Contributing

Contributions are what make the open source community such an amazing place to
learn, inspire, and create. Any contributions you make are **greatly
appreciated**.

If you have a suggestion that would make this better, please fork the repo and
create a pull request. You can also simply open an issue with the tag
"enhancement". Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feat/amazing-feature`)
3. Commit your Changes with
   [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
4. Push to the Branch (`git push origin feat/amazing-feature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the MIT License. See
[LICENSE](https://github.com/zLeki/Goblox/blob/main/LICENSE) for more
information.

<p align="right">(<a href="#top">back to top</a>)</p>



- <https://github.com/zLeki/Goblox>

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->


