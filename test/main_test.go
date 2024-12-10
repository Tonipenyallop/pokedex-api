package tests

// This file is for testing pokemonServiceHelpers

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	pokemonServiceHelper "github.com/Tonipenyallop/pokedex-api/helpers"
	pokemonRepository "github.com/Tonipenyallop/pokedex-api/repository"
	"github.com/Tonipenyallop/pokedex-api/types"
	"github.com/patrickmn/go-cache"
)

func TestGetPokemonsFromCacheByGen(t *testing.T) {

	var pokemonCache = cache.New(24*time.Hour, 24*time.Hour)

	t.Cleanup(func() {
		pokemonCache.Flush()
	})

	genId := "first"

	cacheKey := pokemonServiceHelper.GetGenCacheKey(genId)

	type FakePokemon struct {
		ID   string
		Name string
	}

	pokemons := []pokemonRepository.TmpPokemon{
		{
			ID:   9,
			Name: "Toni",
		},
	}

	pokemonCache.Set(cacheKey, pokemons, 24*time.Hour)

	res := pokemonServiceHelper.GetPokemonsFromCacheByGen(genId, pokemonCache)
	fmt.Println("reflect.DeepEqual(res,pokemons)", reflect.DeepEqual(res, pokemons))
	if !reflect.DeepEqual(res, pokemons) {
		t.Fatalf("expect res: and pokemons: to be same")
	}

}

func TestGetGenIdByPokemonId(t *testing.T) {
	testCases := []struct {
		pokemonId int
		expect    string
	}{
		{151, "first"},
		{152, "second"},
		{300, "third"},
		{400, "fourth"},
		{500, "fifth"},
		{700, "sixth"},
		{750, "seventh"},
		{811, "eighth"},
		{907, "ninth"},
	}

	for _, testCase := range testCases {
		res := pokemonServiceHelper.GetGenIdByPokemonId(testCase.pokemonId)

		if res != testCase.expect {
			t.Fatalf("expect %s to be %s", res, testCase.expect)
		}

	}

}

func TestHelperDescription(t *testing.T) {
	description := `ポケットモンスターのゲーム音楽が楽しめる公式サイト「Pokémon Game Sound Library」にBGMが追加！
同サイトで聴くことができる、『ポケットモンスター ルビー・サファイア』のBGMなどで使用されている全106曲を、動画で公開するよ。

Pokémon Game Sound Library（日本語）
https://soundlibrary.pokemon.co.jp/
※日本語サイトは、日本以外の国からはアクセスできません。

1 タイトルデモ ～ホウエン地方の旅立ち～ 00:12 
2 タイトルデモ2 ～ダブルバトル～ 00:47 
3 タイトル ～メインテーマ～ 01:03 
4 オープニングセレクト 02:45 
5 ミシロタウン 03:59 
6 オダマキ研究所 05:27 
7 ハルカ 06:19 
8 たすけてくれ！ 07:20 
9 戦闘！野生ポケモン 07:43 
10 野生ポケモンに勝利！ 09:13 
11 101番道路 09:36 
12 コトキタウン 10:36 
13 ポケモンセンター 11:41 
14 回復 12:45 
15 視線！たんぱんこぞう 12:50 
16 視線！ミニスカート 13:19 
17 戦闘！トレーナー 13:47 
18 トレーナーに勝利！ 16:06 
19 レベルアップ 16:37 
20 トウカシティ 16:42 
21 連れて行く 17:46 
22 104番道路 18:34 
23 トウカの森 19:45 
24 マグマ団登場！ 20:49 
25 戦闘！アクア・マグマ団 21:34 
26 アクア・マグマ団に勝利！ 24:08 
27 カナズミシティ 24:34 
28 トレーナーズスクール 26:16 
29 海を越えて 27:19 
30 ムロタウン 28:15 
31 視線！うきわガール 30:07 
32 カイナシティ 30:39 
33 海の科学博物館 32:34 
34 110番道路 34:22 
35 サイクリング 35:25 
36 ゲームコーナー 36:51 
37 当たり！ 38:35 
38 残念 38:39 
39 BDタイム 38:44 
40 大当たり！ 39:06 
41 シダケタウン 39:12 
42 113番道路 40:28 
43 ふたごちゃん 42:06 
44 ハジツゲタウン 42:31 
45 ロープウェイ 43:57 
46 えんとつやま 44:11 
47 視線！やまおとこ 46:02 
48 111番道路 46:38 
49 ジム 47:51 
50 戦闘！ジムリーダー 48:51 
51 ジムリーダーに勝利！ 50:37 
52 バッジゲット 51:50 
53 わざマシンゲット 51:57 
54 なみのり 52:02 
55 119番道路 54:15 
56 ヒワマキシティ 56:04 
57 120番道路 56:59 
58 インタビュアー 58:32 
59 サファリゾーン 59:01 
60 視線！ジェントルマン 59:53 
61 ミナモシティ 1:00:36 
62 美術館 1:02:03 
63 わざ忘れ 1:04:12 
64 ユウキ 1:04:17 
65 戦闘！ユウキ・ハルカ 1:05:10 
66 進化 1:06:55 
67 進化おめでとう 1:07:28 
68 フレンドリィショップ 1:07:34 
69 おくりびやま 1:08:50 
70 視線！サイキッカー 1:10:28 
71 視線！オカルトマニア 1:10:59 
72 おくりびやま外壁 1:11:49 
73 アジト 1:13:39 
74 どうぐゲット 1:14:44 
75 アクア団登場！ 1:14:48 
76 戦闘！アクア・マグマ団のリーダー 1:15:38 
77 目覚める超古代ポケモン 1:17:51 
78 日照り 1:18:04 
79 大雨 1:19:10 
80 ダイビング 1:20:05 
81 ルネシティ 1:22:01 
82 めざめのほこら 1:23:34 
83 戦闘！超古代ポケモン 1:24:52 
84 視線！ビキニのおねえさん 1:26:22 
85 サイユウシティ 1:26:53 
86 きのみゲット 1:28:27 
87 コンテストロビー 1:28:31 
88 コンテスト！ 1:29:27 
89 結果発表 1:30:36 
90 コンテスト優勝 1:31:13 
91 おふれのせきしつ 1:31:46 
92 戦闘！レジロック・レジアイス・レジスチル 1:32:51 
93 カラクリ屋敷 1:34:19 
94 すてられぶね 1:35:30 
95 バトルタワー 1:36:42 
96 チャンピオンロード 1:37:52 
97 視線！エリートトレーナー 1:39:13 
98 四天王登場！ 1:40:03 
99 戦闘！四天王 1:40:50 
100 チャンピオンダイゴ 1:42:37 
101 決戦！ダイゴ 1:43:29 
102 ダイゴに勝利！ 1:45:20 
103 栄光の部屋 1:46:14 
104 殿堂入り 1:47:13 
105 エンディング 1:48:23 
106 The END 1:51:20 

#ポケモンゲームサウンドライブラリー #ポケットモンスター #ポケモン

-----
チャンネル登録はこちら
http://www.youtube.com/user/PokemonCoJp?sub_confirmation=1

【ポケモンの最新情報をチェック！】
■オフィシャルサイト
https://www.pokemon.co.jp/
■X
https://twitter.com/pokemon_cojp
■ポケモン情報局
https://twitter.com/poke_times
■LINE
https://line.me/R/ti/p/@pokemon
■Instagram
https://www.instagram.com/pokemon_jpn/
■TikTok
https://www.tiktok.com/@pokemon
-----
©2024 Pokémon. ©1995-2024 Nintendo/Creatures Inc. /GAME FREAK inc.
ポケットモンスター・ポケモン・Pokémonは任天堂・クリーチャーズ・ゲームフリークの登録商標です。`

	

	expected := []types.YoutubeMusic{
		{Name: "1 タイトルデモ ～ホウエン地方の旅立ち～", StartTime: "00:12"},
		{Name: "90 コンテスト優勝", StartTime: "1:31:13"},
		{Name: "106 The END", StartTime: "1:51:20"},
	}


	
	res := pokemonServiceHelper.HelperDescription(description)

	resArr := []types.YoutubeMusic{
		res[0],
		res[89],
		res[105],
	}

	for idx,exp := range expected {
		if resArr[idx].Name != exp.Name {
			t.Fatalf("expect %s to be %s",resArr[0],exp.Name)
		}
		if resArr[idx].StartTime != exp.StartTime {
			t.Fatalf("expect %s to be %s",resArr[0],exp.StartTime)
		}

	}


}
