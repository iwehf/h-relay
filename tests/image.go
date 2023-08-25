package tests

import (
	"h_relay/config"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
)

func CreateImage() *image.RGBA {
	width := 200
	height := 100

	upLeft := image.Point{
		X: 0,
		Y: 0,
	}
	lowRight := image.Point{
		X: width,
		Y: height,
	}

	img := image.NewRGBA(
		image.Rectangle{
			Min: upLeft,
			Max: lowRight,
		})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{R: 100, G: 200, B: 200, A: 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
				// Use zero value.
			}
		}
	}

	return img
}

func CreateImageDataURL() string {
	return "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAABGQAAAH5CAYAAADHtOj/AAAACXBIWXMAAC4jAAAuIwF4pT92AABELklEQVR4nO3dd7ilVXk34N8MvShFUBSwIKIiKvbeu7FX7Bp7b1gSE7HEL4BIEI1iCRYUFQzG3hW7oqIgKhpUEDsd6TDM98eaCQNMOWXvZ7177/u+rnOJcM77POfMnF1+71rPWrJ8+fIAAAAAUGdp7wYAAAAAZo1ABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKCYQAYAAACgmEAGAAAAoJhABgAAAKDY+r0bYPU2PezS3i0AAACwcLdK8oAkd0xywyTbJ9loxX87L8mJSX6V5JtJPp3kN3O98HmPsbZiGixZvnx57x5YDYEMAADAxNkwyVOTvCwthJmPbyfZJ8lnk6z1jbpAZjpYIcNQbZNk1yQ3SHKtJNdMcvUV/7xtkk2TbJb2d3izK3ztuUkuSXJ+klOT/DXJaUlOSXJSkt+mpc+/TXL2mL8PAABgNtw7yTuT7LzAr7/zio9vJHlOkuNH1BcDJZChtyVpyfEdk9wuLYTZNcnWi7jmyoBmiyTbJdltLZ/7pyQ/TfKTFf/7w7TQBgAAYC7WS/LvSV4xouvdLcmPkzwvyQdGdE0GSCBDDzdP20t517QQZjHhy2Jda8XHA1f5dyen7eM8MsnXM4+9nAAAwEzZJMmHkzx8xNfdNMn7k1wvyetGfG0GwgyZgZqyGTKbJLlPkn9ICz526NvOvP06yafSBm19J8myvu0AAAADsH6SI5I8eMx13pBkr1X/hRky00EgM1BTEMgsTXL3JE9O8sgkm3ftZnROT3J4kkOSfDfrGLYFAABMrbcneX5RraelrZhJIpCZFgKZgZrgQGanJM9M8sRM3kqY+fpdkg8lOTjtyDoAAGA2PDLJxwvrXZB2jPYvEoHMtPCnyKjcNcknkpyQ5NWZ/jAmafs5/zXttKZPJrlv2pBiAABgem2Rtjqm0sZJDor3G1NFIMNirJ+2EubotKPZHpbZfIBYkuQhSb6YdjTdc5Js2LUjAABgXF6cdpprtbukve9gSghkWIilSfZI8rO0WSq36NvOoOyS5J1pW5hekjbQGAAAmA4bJHlhx/p7dqzNiAlkmK+HJflxko8kuVHfVgbtmkn+I+0I7ZfHihkAAJgGD0iyTcf6d05y/Y71GSGBDHN18yTfSpsTs3vfVibK1ZLsl7aVaY/M5pYuAACYFg/o3UDGf8w2RQQyrMsWSQ5MWxVz5869TLLrpa0q+kGS23buBQAAWJjb924g7UAVpoBAhjVZkuTJSX6dtkdyvb7tTI3bJPl+knekhV0AAMDk2KV3A0lu2LsBRkMgw+pcO+3EoA8kuXrnXqbRkiTPTdvG9OjOvQAAAHOzZZJNezeR5Lq9G2A0BDJc0ROTHJfkPr0bmQHbJTksyUeTbNW5FwAAYO2GEMYkw+mDRRLIsNJV0macHLLin6nz2CTHJrln70YAAIA1uqB3A0wXgQxJOzXp6LRTgOhjhyRfSbJPzOsBAIAhOq93Aytc2LsBRkMgw+OTfC/Jzr0bIUuSvDLJl5Js27kXAADg8i5I8qfeTST5Y+8GGA2BzOxammTfJB9OsnHnXri8e6YdM36b3o0AAACX8/PeDaQdDsIUEMjMpk2THJ7kFb0bYY12TPKtJI/p3QgAAPB/vtm7gSTf6d0AoyGQmT1bJvlqkkd07oN12yjJx5Ls2bsRAAAgSfI/vRtI8rneDTAaApnZcs20VRe3790I8/LmJG+LYb8AANDbcWnjBXr5WZKfdqzPCAlkZscOaWHMbr0bYUFekOSDEcoAAEBvb+5Y+y0dazNiApnZsF2SI5Ncv3MfLM7j07Ywbdi7EQAAmGGHJ/lBh7rHph3KwpQQyEy/rZN8OcKYafHIJB+PUAYAAHq5NMmzklxYWHNZkmcnuaSwJmMmkJlumyb5bGxTmjYPTlspY/sSAAD0cWySFxXW+6ck3y+sR4H1ezfA2KyXtpxt0gf4/i3JL5P8Isnvkpyc5E9JTk9yZpJzrvD5GyTZPMk2Kz6ukWSntBVCN0wLpzYq6HvcHpbkPUmenmR531YAAGAmvTvJ9ZK8esx13pW+c2sYE4HM9No77U37pPll2rHc30zbl/n7BVzjlLTwZnXWT7JrktsluVuSuyfZfgE1huBpSU5N8srejQAAwIz657TtRK8Z0/UPSjvggym0ZPlyN9eHaNPDLl3Mlz8xySEjaqXCd9LmohyRhQUwi3WTtPDqoUlu06H+Yt0jbWgzAADQx9OTvD3JxiO63iVpK29We6rSeY8xfWQaCGQGahGBzM3SVpaM6oFgXE5JW3r3/iS/6dvK5eyc5MkrPq7TuZe52CvJG3o3AQAAZNe09zh3XuR1fpo2wPeoNX2CQGY6+FOcLldJ8t8Zdhjzy7StNjsm+dcMK4xJkhOSvDZt7sxDk3y9bztr9ZIIYwAAYCh+keSuaavvv72Arz82yZPSVu2vMYxhelghM1ALXCHzgbSVHUN0Ytr+yo+lHRM3SW6RthLlob0bWeHSJM9McnDvRgAAgDW6ftoJqXdJcqMk1007CTdpR2b/McnxSb6V5PNJjpnrha2QmQ4CmYFaQCDz6CSHjaGVxTo3bRXHW9MedCbZLdK+lwd17OHiJE9IcnjHHgAAgIVbkkWelCqQmQ7+FKfDtkne0buJ1fhM2j7KfTP5YUyS/CQt4b5vkuM61D8/bfmjMAYAACaXVREkEchMi7cn2aZ3E6s4N23K+IPT59Skcftykt2TvDzJeUU1z07ygCSfK6oHAADAGAlkJt/9kzymdxOrOCbJzTP9802WJdk/7cjsL4251qlJ7pPkG2OuAwAAQBGBzGTbOG11zFAckuQOGd7JSeN0Yloo9oIkF4zh+n9Mcs+Ysg4AADBVBDKT7cVpk7uH4F/STng6v3cjHSxP8p9JbpXkZyO87u+S3H3E1wQAAGAABDKT6+ppx0j3dklaEPOm3o0MwC+S3D7Jh0dwrV8muWuSE0ZwLQAAAAZm/d4NsGD/nOSqnXu4OMmjknyqcx9Dcl6SJyb5fpIDkqy3gGv8JMn9kpwyurYm0rZJ7pbkjkluluQ6SXZI26qXtC1iZ6QNjv5p2s/ti2nbyGDarJ/k1klummSXJDuv+Ngi7bngKrnsJstZaasV/5620u6EtK2kx6dtfzy1snEASmyVNsdxt7Tnh+sm2T7J1dIO/7jKar7mzCSnp73m/P2Kj1+s+Dgms7nyHUotWb7ciVtDtOlhl67tP2+f9gJ747V90pgJY9btfkkOy/yCs+8k+Ye0N1Sz6CpJnrDi405JlizgGscm+e8k/5U2gwcm1a3STqu7S9p8rk1GdN1fJvlq2uP3kWmP5wBMlm3S5hjeM+010y4jvv4laaHMN9NOGD0yAppBOe8xNrtMA3+Kk2nP9A1jLk3y2Ahj1uWLaas7Tprj5385yX1TG8YcmTYDZxwfW86jj22S7J0WoLwzyZ2zsDAmaatpXp+2UuawJNdKctsxfY9r+thtgb1Pknul9me66yJ6PaCwzy0X0WfSTm77t7QVLT9Kslfai+1RhTFJcuO0QeRfSvKXJG9L+72ZVO9P3Z/vU0u+o4XZMnU/hyNH0O8mSY4r7HnVj2VJ7jGC72Hcjkifn8+pac+dDM81k7wwyXeT/C3tQI2nZfRhTNJWZt4qyUuTfC7JaWk3vB6bZKNFXnvL1P19fv8ie12X1xV+L08d8/dCBwKZybNlkmd07uH5ST7RuYdJ8fO0cOGX6/i8TyR5UNqWp1myfpIXpb35fFVWv5x2Mdd+cJJz0rZp/GSE116X5xfW6mXPwlqfTVs+Pa2WJnlMkm+nvUF9TZKdimpvnRbOHJPka2l3W6HC+Un2yHhOKFyXpUk+lHYzYKiek+ThnWo/NcmfOtXmypak3bA7IsnJSQ5MWzW50BtXC7VJkkck+WhamH9gxhMEwUwRyEyeZyXZvGP9fZIc1LH+JPpD2iyUH63hvx+S5NFJLirraBhukLZF660Z3zykTyQ5e8U/v2NMNVbniRltuDQ0N03tG/d9CmtVWpo2FP34JB9LW3Le0z2SfD4tGLpN516YDccleVmn2tdKcnDq39TOxa5J/qNT7bcn+Uyn2lzeekkenxaYfzEtoFvIbMJx2DJtpc7xST6d5BZdu4EJJpCZLBumrSbo5Qtpd26Zv1OS3Dttpcaq/jPJU9KWT8+ShyQ5Om0r0Tgdsso/fyR128E2T3ujPa1eWVjrqCTfKqxX5f5pw6g/kBZODsmd0n7u700bEgnj9M70W3X74LQVYkOycdoKhB5b03+W5BUd6nJlD0mbiffhtJsgQ7UkbYX30Un+J21LLDAPApnJ8oi0gb49nJQ2aHXWgoNROittyenKUObf014ILu/WUR/PT3vSHvdKr7+kzchY6dy0N79VnldYq9IOafvHq+xbWKvCddP+/n8+w36RnSRPT9sq9oDejTD1npG2mrSHN6edTDMU+6XPY0PPLWRc5kZp20c/mcXNTuvhoWmh3tvS/yRYmBgCmcnytE51L00LY07vVH+arAxlnpp2dPms2TNtOXTFEvFDc+UAsXK73a5pW9WmzUuSbFBU64RMz7yq9dIGI/487UXrpNgubZjjm9PmMsE4nJ62NWOtR0yOyUZpK1I261D7ih6cfjPIXprpntU1dBukDYc9NpMxcHpN1ku72Xh8+s1AgokikJkc10rb8tLD/0ub9cFonJXalRpD8bS0N3VVPriaf/fLjOZ0kLmatlUyW6TNsaqyX/q8QRu1G6TNZdk/yaade1moPdNWnNnCxLh8K8kbO9W+UdppbD1dK8n7OtU+Ism7OtWmbfP5XtqpelU3PMbtmml/r94fq2VgrQQyk+OJ6fPndXzaEaywGHdM7Yu9Y9OG4K1O5XDfh6etMJgWz07dsOJTsvpQbdI8O+2Er9v3bmQE7pH2puE6vRthar0xLbzs4Rlpp531sDRt5tnVOtT+Q5JndqhL88QkP0w7XnoaPSXtOXD3zn3AYAlkJkevAaHPTHJhp9pMh22SfDy1d33W9kb+f9Lmy1TYILUrSsZpwyQvLqx3YNpMg0m1RZLD07bJDWErxKjcMG3F5NAGETMdlqVtXTqjU/13p815qvaqJPfsUNeW9H7WS3ueOyTT9RyxOjulhfnTfNgBLJhAZjLcMMlNOtQ9LP3uVDE9Dkpbulrl0rT5MWtycZL3FPWStFBzGmZvPCFtSX2Fc9NOXplUu6edOPGozn2My/ZJvp7k+r0bYSqdnLZapYct0p4/Kh+zb5fkDYX1VvWmJN/sVHuWbZHks2nHRs+KjdO26/97hnnUPHQjkJkMD+xQ8+K0OzawGI9I8sjiml9K8ud1fM57Undi2A5pgxon2ZIkLy+sd3CS0wrrjdKT0+4E7tS7kTHbPu2kqK17N8JUOiK1Q9hXdYe04aoVrpL6AGil76ZfEDTLtk0LtO/Xu5FOXp3kQ2mrboEIZCbFgzrUPDjJiR3qMj02SfIfHeoeMofPOTnJZ8bdyComfbjvA1O3Sm9Z2vDbSfXWtDuBs+AGST4WryUYj5elnUrWwz8luXtBnYPSJ7w9K21r2CUdas+yHdKCsFv0bqSzx6etwgfiRdQkuEqSuxTXvDjtZCVYjOcnuXZxzb9n7sckVw73vXeSXQrrjdorCmsdFmHwJLl3rKZkPM5PskeSCzrUXpp2F3+cQ3aflPbGtIdnJTmpU+1ZtUPaKY87d+5jKO7WuwEYCoHM8N079UfgfSzJ74trMn1e0qHm4Zn7INgvJ/nNGHu5okldJXOb1L5wqjwandF4XZJdezfBVDoubaVMD9unrRYex7yLnVN7U2BVB8fqhGortymZuwVciUBm+O7UoeYBHWoyfXqcGjCX7UorLU/t4NinJNm0sN6oVK6O+Ura8ZhMlg3T780l0++dmfvKx1F7SEYfpm+Q5CNJNh/xdefi10le1KHuLNs0yediZQywBgKZ4bt9cb2fJPlxcU0YhZOSfGOeX/O+1B3rvmX6LU9fqJ1SO5TZ6pjJdbckD+3dBFPrGUn+0Kn2W5LcbITX+7cktx7h9ebqorQtYOd2qD2r1ktbudvjzxuYEAKZYdsgya2Ka76vuB6MyofSVr3Mx+lJPjqGXtZk0rYtvTx1zxPHpJ2QxeR6XRxnynicnhZoX9qh9kZpzxOjWOF4n9SuOlzVq2MFYrU3pc9JqcAEEcgM2y1Se1rH8rT5MTCJPrjAr6s8WvUWqV/1tlDbJHlaYT2rYybf7knu37sJpta3kryxU+0bZ/HbubdN8oH0CS2/ENvRqz0qBp4DcyCQGbbqJY7fSfK34powCj9I2xu/EN9P8tPRtbJOzy2stRjPSzu6vMLvIwyeFi/o3QBT7Y1Jvt2p9jOTPHqBX7skbQXyNUfXzpz9NW2G2XxXkLJw10vyX72bACaDQGbYblRc75PF9WBUFro6ZqXKgaSPTVt9MmSbJHlhYb39k1xSWI/xuX/a6TQwDsvSti6d0an+u5NcZwFf98Ik/zDiXubqKXGzrdJ6ST6c5Kq9GwEmg0Bm2KoDma8W14NRuCiLX11xaJKzR9DLXGyU5B+Lai3UU1MXGp0RdxKnydK00BHG5eS0Ib89bJn2fLH+PL7m5kn2HUs367Zfki92qj2rXprkDr2bACaHQGbYblhY6/S0oZrQ2/Fpwwdvn2S7JFdLCycfl3b86V+v8PmfS3LaImuem7a3v8pzMtzH36VJXlZY751Jzimsx/j1WgnA7DgitfO/VnXHJHvN8XM3TRsIvNH42lmjo5O8pkPdWbZzkjf0bgKYLPNJ+Km1SZIdC+sdlT6nF8BKf0u7s/SRXHmv++lJfpX2wvaFSe6bNqvigRldkHJQ6rbpXC/JA5J8tqjefDwi7UVlhQuSHFhUa1IsT/LDtAGmP05yQpI/57LQar0k10j7M9o9yd2T3CXDej6/c5LN4nhdxutlaX/3b9Kh9j+nrSo+ch2fd0DqVzsn7XfvcWkrSKlzYOpmry3E2WnPLUcn+U3a88tfkpyX5PwVn7Np2vewXZKdVnzsnuROaYOpgREb0gs4Lm+n1E7i/2FhLbiiHyd5SJI/zeFzlyX5/IqPmyT53xH18Isk30hytxFdb12em2EGMnsW1jokV17xNKt+n+TtabMH1vV7cFra39dPpd2N3TZtTsRLMoz5LRsmuVdafzAu5yfZI+31S+WJlElbSfihtO1Ia1qh+ci0QcA9vCALH3TPwjww7UbL0Pw+bZvd4Wkr4Zet4/PPXPG/v0k77GNVN0zysLRtqbcYWYcw44a6ZJ76Sfw/La4HKx2Vdpd/LmHMFf08o70DWDnc94FpK2WG5C5JbldUa3nafINZ97ckz05y/bSjvxfye3BK2s/yBmmn0AxhQPKdezfATDgutVssV7V9koOz+ptnOyZ5b207/+ejSd7fqfasWprhPZ99I8l9klw3yT+lrYpZVxizLr9Ksk+SWya5adrJYRcv8pow8wQyw1W9LPD44nqQJCemzZsYygyRT6RuxcaStDfiQ/LKwlqfiDu4h6bdcXx3RhOinJ/ktWlLy/88gustxu6d6zM73pn2eNLDQ9JWO65qvbTf7S3Lu0l+lzajjFp7JLlx7yZW+FFaIH73JF/J+I47Py7tgILrpwUzwAIJZIZru8Jay9P2kUKlS9KWvZ7au5FVXJzkPYX1np4+wx5X58ZJHlRYb2h3EytdnHZKzBNy2fLwUToqyW3Tlpz3Yjk7lZ6R5A+dau+ftlpgpX9JnxViy9IeU87qUHuWLU3yut5NpN3YenHagQhX3Go0TienBTN3THJsYV2YGgKZ4bpGYa2/xuA36u2f9sZxaN6dugHX2yR5TFGtdXl5Ya3vJPleYb0hOT9tu9q4j/r+Q9rw69PHXGdNtskw5tkwG05P8vj0OZxgo7RtQpukrU57bYceknby06w+rvb00LTtoj39PMmt0oYKL3Zb0kJ9L+1GwH92qg8TSyAzXNsU1vp9YS1I2pvF1/duYg1OTvLpwnrPK6y1JtdM8uTCevsU1hqS89NWIX2lqN5vkzyxqNbq7NSxNrPnW2kzlHrYNS3MPzR9XlsfmWTvDnWpvZmxOv+dtipmCFuAL0wbKP2EXHZqE7AOApnh2qCwVq87qMyu/dKOWRyqdxbWun36b+94Ueoec45P8pmiWkPz5CRfK675+YzuaPj5skKGam9M8u1OtZ+Y5Nod6p6e5EnptzJilt00bVVULwcleXSGM4dvpUPTBgqf3bsRmAQCmeHaorDWGYW14PT0O31irr6c2vkbPVfJbJ4rD6Ucp/0yviGDQ/amJB/vVPvVSS7oUHeHDjWZbcvSti7N0uuaf0y/+Tmz7hkdax+U9tphqM+n30nbNiuUgXUQyJAkf+/dADPl8CTn9m5iHS5Ne7FT5fHpcyJHkjwrdQHwX5IcUlRrSI5K36GPf0ntsOqVrtWhJpycvm+UK70zySd7NzGjNki/LaEfybDDmJV+kHaS5oW9G4EhE8gMV+UKGfs8qXR47wbm6H2pexGxaZKnFNVa1fpJXlJY74DM3gDxZWmnaY3iWOvFeEeHmlt3qAlJckRqQ/Uefp7+80tm2b3S5zHuqCRPy/DDmJW+nbaKC1gDgcxwLendAIzB2WnDByfBaUk+Vljvuan/vX9skh2Lap2T5F1FtYbk7UmO691E2uyeHxXXHMqR7syml6WFFtPowiR7xA21nh7doebf0563J23FyaEZ/8mCMLEEMkClb2WyBg9W3mG9YZJ7FtZLklcW1npXkjML6w3BuUn+rXcTq6je2rBJcT1Y1flpoUWP+Unj9rIMI+idVUvStuJUe0GSEzvUHYWXJPnf3k3AEAlkSNp2CajQ6/SLhfpekmMK6z2/sNZ9k9ysqNbFaduVZs07kpzau4lVHFlcTyBDb8elhRfT5FPpswWRy9wiyTWKa34+yQeLa47SOZmd2U4wLwKZ4TqrsNaGhbWYbd/t3cACVL7wfUjqTqZ5RVGdJPloZu8UkOUZ3pumH6V2hZotSwzBO5N8oncTI/KnmMcxBPctrrcsyZ7FNcfhm5me30UYGYHMcFUO67JChirH9m5gAT6cupPI1kvyzII6uye5d0Gdld5cWGsovpDhLS2/IMkJvZuADp6RyQ+FL0071ee03o2QuxTX+68kvyiuOS6vzGRtXYexE8gM1zmFtbYrrMXs+lMmc4bIuUk+UFjvWWnHaY5T5eqYzyf5WWG9oRjq0vLjezcAHZye5PFpocak2ifJ13s3QZYmuVNhveVJ9i6sN24npJ2CBqwgkBmuMwtrbVtYi9k1yQMI31lYa7skDxvj9a+TdkpDlVlcHXNRWhA1RH/q3QB08q0kb+zdxAJ9P8lrezdBkmTnJFsU1vt0kt8V1qtwYO8GYEgEMsN1ZmGtqxfWYnb9tncDi/CLJN8orDfO4b4vSdsaVeFHmc07ut9I7Ryw+Zj0bRuwGG/M5A2XPzvJE5Jc0rsRkrQtv5WGNotsFL6d6T2SHuZNIDNcZxbW2jq1aT+z6aTeDSxS5SqZuyXZdQzX3So1M2pW2rew1pB8uXcDa/HX3g1AR8vSti6d0buReXhuJvuGxrTZvbDWGUm+Vliv0sd7NwBDIZAZrr8U19uluB6z5/e9G1ikT6T2zezzxnDN5ybZbAzXXZ3fZnb3iVeuppovW5aYdSdnco7f/UCSQ3s3weXcsLDWZ5NcXFivktOWYAWBzHBVv2gWyDBukx7IXJTkvYX1npzkKiO83kZJXjTC663L/pnNkxSWZdinif25dwMwAEckOah3E+twQpIX9G6CK9m5sNZnCmtVOyaejyCJQGbI/lhc76bF9Zg9f+vdwAi8O3WndFwlbW7AqDwpyTVGeL21OS3JwUW1hub3acdLD9UkbdWAcXpZhjvH4uIkj0vtiZvMzU6Ftb5bWKuHH/RuAIZAIDNcJxfXu21xPWbP6b0bGIHfp/aO1ai2LS1NsueIrjUXb0tyfmG9IRn6aRhDHTYM1c5PskeGGaC+Jm0oOsOy+YqPCn9N/XuBagIZiEBmyM5L7VK+28TfB8ZnedqqiWlQOdz3pknuPILr/EPq9r2fn+Q/i2oN0Sm9G1iHc3s3AANyXNpKmSH5cpL9ejfBam1fWOvHhbV6OaZ3AzAE3oAP268Ka22e5OaF9Zgtp6WFMtPgS6k98WIUR2C/agTXmKuDk5xaWG9ohh48XhzH58Kq3pnhDBg9JW1+2LQ8X06bqxfWGvpqy1GYhe8R1kkgM2zHF9e7b3E9ZsffezcwQpemdhjkI7O4F4F3SHKnEfWyLsvShvnOskmYlWSVDFzeM5L8oXcTSZ6a+lM2mbstC2tN+kEIc3FS7wZgCAQyw1a5QiZJ7l9cj9lxdu8GRux9SS4sqrVBkmcu4usrZ8cckdrVQ0M0CauDzJFhsaYpZE/ajLPHp25o++q8NcnnOtZn3bYorDULYcX58XwEApmBO7q43p1Sm/4zO87r3cCInZrksMJ6z0qy3gK+7gZJHj7iXtZm38JaQzUJb1Sn7fdxVgdI9zSNR9r/MH1XJWzYsTZzs0lhrWl7nF6TM3s3AL0JZIbtx6m9W7NBkkcX1mN2VK0mqVS5benaSR60gK97eZIlI+5lTb4Wp4Ikk7EabNp+H6ft+6GP/ZNct2P95yZ5WMf6rNtmhbUmIdwfhSGecgalBDLDdm7q58g8vrges2Eal6R+N7UnBMz3COyrJ3nKOBpZA6eCNMKBej23mcyqaVuV9NC0QKS3g5Ps0LsJBmHafsfWRCDDzBPIDN9RxfXulr53iGCSVB6BfZ8kO8/j81+YZOMx9XJFxyX5QlGtoZuEF9HT9gK4clXSRoW1hmyagsftk/xX7yZW2CrJoVnYFlWmy6xsYataxQuDJZAZvq8V11uS+d+Jh3WZtjeAK304dcuK5/O7uVlq7/buG8e0rjQJJxhN6+9jhcoZEvO1ZWGtSfh7PhdLk3woydV6N7KKuyR5Te8m6G5Wwt/KQckwSAKZ4ft6h5rPTO0+WabftL4BPCfJBwvrPTVze0P4tNS9wfhDko8W1ZoEF/duYAad07uBGTQtf89fneTuvZtYjb2S3Ll3E1xJ5WuZzQtr9TQrwROskUBm+P6Q5DfFNbdM8vTimjCpKrctbZVkj3V8zvppw3yrHJDpeXM2CtM4L2noKreJWV7fTMPA0dsneX3vJtZgadrWpa16N8LlVD7WbF1Yq5clmY3vE9ZKIDMZPteh5j9l2EuzYSh+nuSbhfXWtW3pEambA3VWkncV1ZoUtm7Vq1whY3l9M+mrHq+aFnis37uRtdgxyXt7N8HlnFlY67qFtXq5RmZnVg6skUBmMnyqQ83t0oaCAutWuUrm1klus5b//sqqRtK+b9tFJs80DWRNalclDXl5/VULa51ZWGscDkpyvd5NzMEjkjy7dxP8nzMLa127sFYv1+ndAAyBQGYyfCN9lsG/Osk2HerCpDkiyd8K6z1/Df/+HkluVdTDxUkOLKrFaE3CSVDzURkKDnkLSWVYNMlb856c5HG9m5iHA5LcpHcTJEn+Wlhrl8Javczn5EiYWgKZyXBxkk93qLtVkr071GX6DHlZ+ChclNql5Y/N6vddv6Kwhw8m+XNhPViTyhN/hhzIWCGzbjsneUfvJuZp47TB6Rv3boT8sbDWLZJsUFivh1v3bgCGQCAzOT7Uqe7Tk9yxU22mxyycFvDuJJcW1do4yT9e4d/tluQBRfWXJ9mvqBasy+mFtQQyzZmFtUZlg7RgYxJPkdwtyf69myB/T91A642T3LyoVi+37d0ADIFAZnJ8JclfOtV+XybzBcxQbZvJWi7N3JyU5LOF9Z6Tyz+GV66O+UyS4wvrTZJJ3soxqU4rrDXkQKZy4PCZhbVG5U2p29I5Ds9N8rDeTVB68umdCmtV2zjJLXs3AUMgkJkcy5J8uFPtXZK8pVPtabNtkq+lne7w3M69MHqVw32vn+S+K/55+9SGfPsU1po0TlmqV7lCZshHtFYGMqcU1hqF+6Q2tB6Xg5Ps0LuJGVd5M+IRhbWq3Tu24UESgcykeXfH2s/OdD8xVNghyffSlh4nbR/7E/u1wxh8McnvCuutDPVekrq95t9L8p2iWjAXl6YulNm+qM5CVIVFy1I7xHyxtk2beTUNtkq7obNe70Zm2C8La905ydUL61V6WO8GYCgEMpPl12lv+Hr5YC4LE5ifGyT5ZtqqhlV9IMkT6tthTC5NO061yoOS3CzJswprvrmwFsxV1YqNjZJco6jWfFUdk/u3TM5KsCVJ3p9ku859jNJdkrymdxMz7CeFtZamDfGfNhsmeWjvJmAoBDKT520da2+W5FNxFPZ83TrJd5NcbzX/bWnai8WHVDbEWB2cdupShaVp81yqhnn+Oskni2rBfPyhsNZOhbXm4zpFdf5UVGcUXpTkgb2bGIO90lZPUO/o4novSAsWp8lj470E/B+BzOT5fGqXS17R9dIGl1buVZ9k901yZNb+xLN+ko+nrXZg8p2a5LDCejsW1tovdSdJwXxUHkd748Ja83Hdojq9DhiYr5sn2bd3E2OyNG3r0pCHTE+rP6YN8a+yS5L7F9ar8ILeDcCQCGQmz6VpJwX0dNsk/5Nk0859DN1zknwuczuhaoO0N/H3HGtHVKkc7lvlr0kO6d0ErEHlqo0hHkW7fuqGvZ5YVGcxNks74nrD3o2M0Y5J3tu7iRl1ZHG912V6VsncK467hssRyEymj6X22L3VuXvaVomrdO5jiNZLckDam/L5DN7bJO1nepcx9ESt7yY5tncTI3Zgkgt6NwFrcGJhrdsX1pqr3VI32Lv364+5+I8kN+rdRIFHpB26QK0ji+vdNskexTXHYWnaSltgFQKZyXRJkjf2biLJPdKOcN62dyMDcrW0VTEvXuDXrwxlbjeyjuhlmlbJnJvp+n6YPv9bWOsWGd623cp5IicU1lqIRyV5Zu8mCh2Q5Ca9m5gxX0j9YOu9k2xeXHPUnpJk995NwNAIZCbXIRnGHfhbpx2Bu2vvRgbg1mnD3u67yOtcNcmXMsxl8czdh5P8vXcTI/KeJGf0bgLWojKQ2SDDm+nwsMJavy2sNV/XTnu86uH76XP61MZp27M27lB7Vv0lyY+Ka147beXXpNo2VsfAaglkJtelSV7eu4kVbpDkB0ke3ruRTpakDSj7dkZ37OhVk3wlyU1HdD3q/T3TMXNlWSb7RSCz4Q9Jzius96TCWuuyU+rmjy3LcFfIrJcWhG/ZofaJaSHd3h1qJ23L2v6das+qIzrUfEYm81TOJUk+mGTr3o3AEAlkJttX0k48GoLN056c9s9s3aW5etoWo7cl2WjE194myVeT7Dzi61JnGrb5fDTJ73s3AeuwPMnPCus9IMN5bH5R6gZ+/iLJhUW15utf0uco6GVJnpDkrCSvTVsp08NzU7tSatZ9pFPdD6TdCJ0k/5rhrSqEwRDITL4XJDm/dxOreGmSH6ftsZ92eyT5eZIHjrHGtmnD44bywp/5OS7Jt3o3sUhv7t0AzNGPC2stTfKqwnprsn1qh7r+pLDWfNw5LQzp4XVpg9yTNuPv8UnO7tTLwak7bWvWnZTkmx3qbpl2M3ZS5jfukeT1vZuAIRPITL4T014MDMmuSX6Ytlrmqp17GYcdk3w67e7INgX1tk/y5XiRNakmeZXMl5Ic07sJmKPKQCZJ/jHJzYprXtE+qV2VWv0znout0rYq9XhN+80k/36Ff/e7JM/q0EvSfhaHZn4nPLJw7+5U9wZpB0gMbbj4FT0kbasSsBYCmemwf5Kf9m7iCtZLWy3zq7Sp6tPwd22ztPDrV0keVFz7umkrZYQyk+e/k5zSu4kFsjqGSVJ9t3pp2oqE9YvrrvQPaVtlKh1VXG8u3pXRzW+bjzPSfv7LVvPfPpbkv2rb+T93SfKaTrVnzcfT7/n91hn2SaePSPv5bNC7ERi6aXiTTFsi+8QMc1/3dknen7Z1Y4/U7XMfpQ2TPCfJr5PslXY0dQ/XT/L5DPfJl9W7KMl7ezexAEenzamCSXFC2uqESrdKn5NDrpv6O89np61+HZJnJnl0p9pPTxsmvSYvTnJ8US9XtFf6zNNZqK2SPC3Jh9JmQZ2bNhdqedq2/J+nzTN7ZmpWJs/VhUne0bH+LZN8L22o85C8NMnhEcbAnAhkpsfPM4z97Gty47QtPj9LW8q7ad925mTLtJOsTkrbdnKtrt00u6XdEbla70aYl3elz3Goi+F4SibR5zvUfHGSFxbW2zbt+6w+seTrWf1qkF5ulOSATrXfleQT6/icc9NuRPW4WbY0bevSVh1qz8f10m5Y/CVttdkT0l7nrPoaceO0rfCPTdsi9Oe0MHKX0k7X7O1pf9a9XD9tkPRjO/aw0lXSgpj94z0mzJlflulyYJIv9m5iHW6S9kLm5LQ3fLt37ebKliS5Y5L3JPljWo/bde3oynbLZOwd5jInZTgnos3FiWkvqmDSHNap7oFJ9iyoc+0k30gLI6p9uUPNNdko7SZPj5s7v0xbATAXxyR5xRh7WZsdM9zVmSuHYv8ibaXRhvP42vXTjp3/Wdo28t7zck5N/1lxm6WtIPpg+q0gul/a3/dHdaoPE0sgM12Wp91dOKl3I3Owddrqk5+kre7517Sllz22NK2XtrR37yS/TfKdJM/IsFfx3DZt4OqQe+Tyer9gm4/907ZCwqT5Vta+jWSc3px2JO3mY7r+/ZP8KG3FabVLMqyQdp/0uaFzYdqql/mcbvn2tIMAenhEak/hmovN034ee2dxA6k3TNua9dX0v0G1d9qx5709KW2b3LMyv5BrMXZOC8K/kLbiCZgngcz0OS0tnR7iPJk12TXJG9JOb1i5FPW5aUdnj2NY4sZpgcZLc9nA1W+l3a257hjqjctt02/CP/P3hdTPt1iI09OWjsMkujRtq0YvT06bmfbwjO4Gw7WTHJK+M8S+kORvnWpf0QPTton1sGeSY+f5NcvTTuT60+jbmZMD0lYnD8EWaQHKA0d4zbulvYbruZX7tLSQcAiulrYS/YS0+YebjanO7mmr1I5PvzlOMBV6nQzAeP0o7Y7I+zv3sRDXSEv4n7Ti/1+Q9mD/67TTjU5Oe1FzalqQcmGuvHd3k7Qn/a3Tnph2TAtadkp7UXKD9F/iOgq/SfLPvZtgzi5Ne5G0d+9G1uEd6bsfHhbr4LRtIr2GyF8nyRFpg7EPTAv+z5nnNZYkuX3ac/nj03845lBC2pUHBfTwmST/ucCvPTXt8IWvpP5m6MZp21luk/aaqpcN0ubu3HYM175p2p/PPdLve/yPtOBt5071r2jHtJW5+6WddvSRtODqvEVc88ZJHprLZv0AIyCQmV4fSHtS+JfejSzSxmkp/O592xic45LcO8lfezfCvBycthqsainxfF2Q9gYSJtmv0uZs/UPnPm6ZFh4clHYk93fSVlf8Nu2xe+VK1qVpK192SLtpcNsk98pw5pcdn+STvZtIC6k+mD6rhP6cdgrQYoazfz0tkO9xI2W3tK2oz+tQe6X/lxaYjMvt01YDPWeMNdbmgrTh3j0Gi6/NZkmesuLj4rSj67+f9jh5Qtq8xPNyWWi8yYqPa6QNDL5+kpsluWuGdcIVTA2BzHR7bdrKkCd27oPROipteNqZnftg/k5Jm8PwhN6NrMH703qESXdA+gcyK22c5L4rPibRPmkr/HrbM8l9OtRdnrYV7dQRXGuvtFDiDiO41nw9N2323P90qH3rtLmB4/bstJUg3yiotTpfSNsy+fhO9ddlgyR3WvEBDIQZMtNtedr0+qGfvMTcfSnJPSOMmWQH9W5gDZYneUvvJmBEvpI2l4zF+V36zuRZ6VZJ3tSp9r5pf59G4ZK0N+u9BsAenLYSq9p+qdtC2Pt57AXpNy8ImEACmel3UdqU/a/3boRFe1+SB8V8j0n37bTjOofmiLTlyzAtXt27gSnwqrTXET1tnjYDpcccnR+mnQI5SiemnYLTw1ZJPpzaOXq3Shu8W1lvnFuj1uWMLH57GzBDBDKz4by0N/JH9W6EBdsrbVjcxb0bYSSGeAT2vr0bgBH7Ska3smEWfTdtGGhvb0+fQannJHlcxvO8e1iS947hunNx19TOsXlKYa2Vem/V/1LazByAdRLIzI7z0vavWykzWS5O8tS0QbBMjw9l/qeujNO3IrBlOu2Ztk2E+VmW5EXpf5f/cenzhj5pA3B/M8brvyRtYHIPe6Vujsj9iuqs6v4dal7RXhEIA3MgkJktZ6WtlPlM70aYk7+lzYv5QO9GGLm/JzmkdxOr2Kd3AzAmxyR5c+8mJtA+6T+DZ6f0m7l1aMb/GH1ukj1y2WlbldZL+x63GnOdrZPsMuYaq3OtJNfuUHdVy5I8JsmvO/cBDJxAZvacl+SRaXfoGa6j004l+HbvRhiboWxb+kXaEcEwrd4Qb4rm41dJ3ti5h/XTZp1ctUPt36bu6ORjkryiqNYVXTvJu8dc4/pjvv5Qa690RpIHZzQndAFTSiAzmy5KO8Lx9b0bYbUOTXLnJCf3boSx+lmGEbi9Of23JcA4XZB21Hzv4bST4Pwkj077mfX0+iS371B35SlIfy+s+fYkny6st6pHZbwDhrcd47XX5Woda6/q10kemGFtUwYGRCAzu5YneV3ai9Qey2W5sgvT9qw/Ie1FMdOv9yqZP2UYR9rCuP0obWYHa/fc9D8F7h7pd0LWa5P8oLjm8rSh/b2OSj4gya5juvaGY7ruXGzesfYV/TDJQ+O1HbAaAhkOTRvs9rvejcy4E9LuBvZ+g06tjyc5pWP9A2LVALPjnUk+2LuJATsw/WeWXS1tdkuP16dfT795WqemnQx0aYfam6QdK77xGK7d84Zf5Sqnufha2hxHoQxwOQIZkja471bpt2R21n00yS2T/LRzH9S7KMl/dap9dsY/PwCG5plpR9JyeR9P8rLeTaQ9Hm7foe5p6ReIrPT1JHt3qn3TJG8Zw3X/NoZrztUQ57asDGXO7t3IAJzbuwEYCoEMK52RtpzyxZHeVzk1bQL/4zK8OznUeVf6zHB5V9rJazBLLkryiLQtBDRfT9squ6xzH89Lex3SQ88tQ6vaK8n3OtUex8//hBFfbz6GOsj7a2knaA4xMKry+yT36d0EDIVAhlUtT1uyfMskR3XuZdp9Mu2O1OG9G6G7E1N/ytHFSd5aXBOG4twkD4jnuST5bNod+95bF3fLeFZozMU7knyqU+0ruiTtJs2ZneofnGSHEV7vrCQ/H+H15uqkJH/uUHeufpzktmmnHM6ao9K26P+ydyMwFAIZVuf4tLkyr4glhaP2l7QTHB624p8hqZ8d9KEkfyyuCUNyWpL7pt2tnlWHp60WOq9zHxsn+UjGM8NkXY5L8vIOddfmpIz35KO12Trt+WG9EV6z+oZDkny+Q835+l2SO6SForPikCR3y7DDMignkGFNLkmyX5IbJzmicy/TYFmStyW5YdoLT1jV/xbX63UnGobkrLSVMr3mOPX0+iSPTf+VMUl7PNqtQ90LkuyR/kd8r87hSd7TqfbdkvzzCK93yAivNVe9h1PP1dlJHpzkn9J/y+A4XZh2gtuTM8zfN+hKIMO6nJzkkWl7PY/u3MukOjJtaPKLYpAbq1d5h/az6bOEHIbooiTPSPL09D0RpspZaatiXpc+s6uu6KFps0t6eFmG/Vj4kvTb1rFX2krpUfhZki+O6Fpz8e0k3y+st1jL04Y53y19Z+6Myy/Ttigd1LsRGCqBDHP1lSS3TjuF4Dede5kUP01y/yT3SHJM31YYsO2SPKWw3r6FtWBSHJzkjpnumQ5fS5td9onejaywffqtTvpk6reKztd5aSt4egSF6yU5NMlWI7renqlZAbJ8Ra1J9J0kt0jynxlGWLpYlybZP04RhXUSyDAfy5N8OMmN0oKZaX7huhjHpg3lu2Vq7woxmV6cZKOiWkcl+WZRLZg0R6etZnxThrGVZ1TOSPL8JPdOW/U6BEvTtrJcrUPtP6atiJoEx6bfjJtrJ3n3iK51XJI3juhaa7N/kh8U1BmXc5K8IG3g748697IYR6etinl5bFGCdRLIsBCXpAUzu6UtN/5K33YG46tJ7pdk9yQfzXTc4WC8rpq2r7rKmwtrwSS6IMm/JLlJ2ryvSX4cX5Z2t33ntJOEhvS9vDpt9Wi1S9NuKJ3WofZC9TwF6lEZ3YDhf8t4B/x+PW0WyzT4UZLbpf1d/V3nXubjL0memeQ2SX7YuReYGAIZFmN52ouE+6QNq31rktO7dlTvtLTv+2Zpdx+/lGG96GXYnpNki6JaJ8SAbpirE9JOxNs97ffm0q7dzM8FaW/ib5B2t31oz8u3Txsq3MPeaXPdJsnyJP+YfifjHZBk1xFcZ1mSR2c8qzSPSvLwJBeP4dq9XJrLVqU/I/XD/+fjL2lh2A2SvDeT9XgJ3QlkGJVfpw2g2y5tYvzHMr3LFC9Ie4G+R9oe+JekDa2D+dgwyUsL670lXiTBfB2bNtj+umlbmf7atZu1+3WS1yS5TtoWpSHeWb9q2pvM9TvU/n7asNpJdFraaokej+GbpK36HcWx5OelnWw2ypsDn0+7IXbWCK85JBelzVq6UZKHpH2/Q7nx97O0FVQ7pYWd5/RtByaTQIZRuzjJZ9LCiqsleVjaE8mQX8TOxWlpL0gem2TbtBfoH8tsnMrBeDwlLcCscEom5xhQGKKT07Yy7Zj25u8dSf7ctaPm12mrNO+QtlL1/yX5W9eO1u6gtDdv1c5O8oS0LdeT6si0P98ebpoW6o/CeWlboV664p8X6sK0VRkPSvL3EfQ1dJcm+XSSB6aFrq9I8uMOffwxydvStiXdLO149vM79AFTY8ny5UMJWVnVpodN3Y3spWkP3HdPcpckd02yTc+G1uGMtMFwX0+bkfPTTOfqgs1Tc6fy4iTnFtSZFEuTHJ+2vLfCa1MzUHGoNknd4OSzMpy7l2uyWZINCupckum+Y7ok7XntTmnPa3dMG4Q6LsvTHjeOSlvt8aUkvx1jvXHYslPdaXkOWpq2yqiH5Rn9KpQd01Z1PS1t1ehcXJx2AtQbMnl//8fhmmmrju6R9lh0vRFf/9y0x5sj0w6q+FFG8xy3JHVbti/K4sK/ddk4o1lBNhfnZZWh8+c9xtqKaSCQGagpDGRWZ6e0Ey1ulbZP/4Zpqf+S4j5OSnuR+4u0uw1HZdh7dZl8j0pyeFGtc9N+ryZpiCVMqi3TnstunGSXtG2t10hy9bQVcZumhWGbXuHrzk97kX122orSU9NWuvw2bZ7Nr5P8asV/h2mzRZJHJLlXkpunvT5c+TtyQdrvwc/SDk84Ip7P1uYaaSuabprk+mnbLbdPuwm6da782LMsyZlpNyL/lOQPSX6T9rr4mBX/W3FkOQsgkJkOApmBmpFAZnU2SnsRu1OSa6XdPblW2gvZLdJe7G6Zdsd7s6x+dceFaU/g5+eyJ5kz057A/5Dk9yv+9+S04GUa7poxWX6Y5NZFtd6e5IVFtYC5W5Lhr6aCnvyOjI+f7RQQyEyHHkPVYG0uTLsLMt8huZ5YmBT3Sl0Ysyyj2/cPjJbnLFg7vyPj42cLAyFWY1p4YmFSvKqw1uFJTiysBwAAzJFABqDOLZPcp7DemwtrAQAA8yCQAajzysJaX01ydGE9AABgHgQyADWun3a6UpV9C2sBAADzJJABqLFnkvWKah2b5EtFtQAAgAUQyACM3zWSPLWwntUxAAAwcAIZgPF7cZKNi2qdnORjRbUAAIAFEsgAjNdVkjyvsN7+SS4prAcAACyAQAZgvJ6TZIuiWmcmeW9RLQAAYBEEMgDjs2GSlxbWe0eScwrrAQAACySQARifJye5ZlGtC5McWFQLAABYJIEMwHgsTfKKwnofTPLXwnoAAMAiCGQAxuNhSXYpqrU8yVuKagEAACMgkAEYj1cV1vqfJL8qrAcAACySQAZg9O6R5LaF9fYrrAUAAIyAQAZg9CpXx3wnyXcL6wEAACMgkAEYrd2T3K+w3j6FtQAAgBERyACMVuXqmOOTfLawHgAAMCICGYDR2SnJowvrvSXJpYX1AACAERHIAIzOnknWK6r1lyQfLKoFAACMmEAGYDSunuRphfXemuSiwnoAAMAICWQARuPFSTYuqnVOkoOKagEAAGMgkAFYvM2TPK+w3ruTnFlYDwAAGDGBDMDiPTvJlkW1Lk5yQFEtAABgTAQyAIuzYZKXFtb7WJKTC+sBAABjIJABWJwnJtm+sN6+hbUAAIAxEcgALNzSJK8orPf5JD8rrAcAAIyJQAZg4R6a5EaF9fYrrAUAAIyRQAZg4V5VWOvHSb5WWA8AABgjgQzAwtwtye0K65kdAwAAU0QgA7Awry6s9bsk/11YDwAAGDOBDMD83TzJ/QvrvSXJssJ6AADAmAlkAObvlYW1TktycGE9AACggEAGYH6um+SxhfXenuT8wnoAAEABgQzA/OyZZL2iWhekBTIAAMCUWbJ8+fLePQAAAADMFCtkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIoJZAAAAACKCWQAAAAAiglkAAAAAIr9f0UNrpIrb4gAAAAAAElFTkSuQmCC"
}

func prepareResultImagesForTask(taskIdStr string, numImages int) error {
	appConfig := config.GetConfig()

	imageFolder := filepath.Join(
		appConfig.DataDir.InferenceTasks,
		taskIdStr,
		"results",
	)

	if err := os.MkdirAll(imageFolder, os.ModeDir); err != nil {
		return err
	}

	for i := 0; i < numImages; i++ {
		filename := filepath.Join(imageFolder, strconv.Itoa(i)+".png")
		img := CreateImage()
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return err
		}

		if err := png.Encode(f, img); err != nil {
			return err
		}

		if err := f.Close(); err != nil {
			return err
		}
	}
	return nil
}