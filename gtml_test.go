package gtml_test

import (
	"gtml"
	"testing"
)

func TestMain(t *testing.T) {

	_, err := gtml.NewGtmlRootFromStr(`
		<div _component="GuestList">
			<div _for="guest of guests []Guest">
				<h1>{{ guest.Name }}</h1>
				<p>The guest has brought the following items:</p>
				<div _for="item of guest.Items []Item">
					<p>{{ item.Name }}</p>
					<p>{{ item.Price }}</p>
				</div>
			</div>
			<div _for="color of colors []string">
				<p>{{ color }}</p>
				<p>{{ color }}</p>
			</div>
			<div>
				<p _if="loggedIn">Im logged in</p>
				<p _else="loggedIn">In not logged in</p>
			</div>
		</div>
	`)
	if err != nil {
		panic(err)
	}

}
