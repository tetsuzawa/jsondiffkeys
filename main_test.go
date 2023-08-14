package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_jsonDiffKeys(t *testing.T) {
	type args struct {
		a []byte
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				a: []byte(`{"userId":1137736635576750100,"viewerId":"33","sessionId":"3730e3fe-687b-41b3-8bdf-92fbe138ad7e","createdAt":1661606352,"updatedResources":{"now":1661606352,"user":{"id":1137736635576750100,"isuCoin":6500,"lastGetRewardAt":1661606352,"lastActivatedAt":1661606352,"registeredAt":1661606352,"createdAt":1661606352,"updatedAt":1661606352},"userDevice":{"id":1137736635580944400,"userId":1137736635576750100,"platformId":"33","platformType":1,"createdAt":1661606352,"updatedAt":1661606352},"userCards":[{"id":1137736635580944400,"userId":1137736635576750100,"cardId":2,"amountPerSec":1,"level":1,"totalExp":0,"createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635580944400,"userId":1137736635576750100,"cardId":2,"amountPerSec":1,"level":1,"totalExp":0,"createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635580944400,"userId":1137736635576750100,"cardId":2,"amountPerSec":1,"level":1,"totalExp":0,"createdAt":1661606352,"updatedAt":1661606352}],"userDecks":[{"id":1137736635580944400,"userId":1137736635576750100,"cardId1":1137736635580944400,"cardId2":1137736635580944400,"cardId3":1137736635580944400,"createdAt":1661606352,"updatedAt":1661606352}],"userLoginBonuses":[{"id":1137736635585138700,"userId":1137736635576750100,"loginBonusId":1,"lastRewardSequence":1,"loopCount":1,"createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635589333000,"userId":1137736635576750100,"loginBonusId":2,"lastRewardSequence":1,"loopCount":1,"createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635589333000,"userId":1137736635576750100,"loginBonusId":4,"lastRewardSequence":1,"loopCount":1,"createdAt":1661606352,"updatedAt":1661606352}],"userPresents":[{"id":1137736635593527300,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":3000,"presentMessage":"リリース記念プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635593527300,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":1000,"presentMessage":"１ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635593527300,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":1000,"presentMessage":"２ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635593527300,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"３ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635593527300,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"４ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635593527300,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"５ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635593527300,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":3000,"presentMessage":"ハーフアニバーサリープレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635593527300,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"７ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635593527300,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"８ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635597721600,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"９ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635597721600,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１０ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635597721600,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１１ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635597721600,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":3000,"presentMessage":"１周年突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635597721600,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":6000,"presentMessage":"周年記念プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635597721600,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１３ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635597721600,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１４ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635597721600,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１５ヶ月突破プレゼントです１","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635597721600,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１６ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635597721600,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１７ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635597721600,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１８ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635597721600,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１９ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635597721600,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"２０ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635601915900,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"２１ヶ月プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635601915900,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"２２ヶ月プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635601915900,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"２３ヶ月プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635601915900,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":6000,"presentMessage":"２周年記念プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1137736635601915900,"userId":1137736635576750100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":6000,"presentMessage":"２.５周年プレゼントです！","createdAt":1661606352,"updatedAt":1661606352}]}}`),
				b: []byte(`{"userId":1140558566977966100,"viewerId":"33","sessionId":"0c491546-151d-41c6-be16-06071850646b","createdAt":1661606352,"updatedResources":{"now":1661606352,"user":{"id":1140558566977966100,"isuCoin":6500,"lastGetRewardAt":1661606352,"lastActivatedAt":1661606352,"registeredAt":1661606352,"createdAt":1661606352,"updatedAt":1661606352},"userDevice":{"id":1140558567003131900,"userId":1140558566977966100,"platformId":"33","platformType":1,"createdAt":1661606352,"updatedAt":1661606352},"userCards":[{"id":1140558567003131900,"userId":1140558566977966100,"cardId":2,"amountPerSec":1,"level":1,"totalExp":0,"createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567003131900,"userId":1140558566977966100,"cardId":2,"amountPerSec":1,"level":1,"totalExp":0,"createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567003131900,"userId":1140558566977966100,"cardId":2,"amountPerSec":1,"level":1,"totalExp":0,"createdAt":1661606352,"updatedAt":1661606352}],"userDecks":[{"id":1140558567003131900,"userId":1140558566977966100,"cardId1":1140558567003131900,"cardId2":1140558567003131900,"cardId3":1140558567003131900,"createdAt":1661606352,"updatedAt":1661606352}],"userLoginBonuses":[{"id":1140558567007326200,"userId":1140558566977966100,"loginBonusId":1,"lastRewardSequence":1,"loopCount":1,"createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567011520500,"userId":1140558566977966100,"loginBonusId":2,"lastRewardSequence":1,"loopCount":1,"createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567011520500,"userId":1140558566977966100,"loginBonusId":4,"lastRewardSequence":1,"loopCount":1,"createdAt":1661606352,"updatedAt":1661606352}],"userPresents":[{"id":1140558567015714800,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":3000,"presentMessage":"リリース記念プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567015714800,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":1000,"presentMessage":"１ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567015714800,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":1000,"presentMessage":"２ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567015714800,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"３ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567015714800,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"４ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567015714800,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"５ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567015714800,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":3000,"presentMessage":"ハーフアニバーサリープレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567015714800,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"７ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567015714800,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"８ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567015714800,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"９ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567019909100,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１０ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567019909100,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１１ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567019909100,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":3000,"presentMessage":"１周年突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567019909100,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":6000,"presentMessage":"周年記念プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567019909100,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１３ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567019909100,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１４ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567019909100,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１５ヶ月突破プレゼントです１","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567019909100,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１６ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567019909100,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１７ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567019909100,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１８ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567019909100,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"１９ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567019909100,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"２０ヶ月突破プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567019909100,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"２１ヶ月プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567019909100,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"２２ヶ月プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567024103400,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":2000,"presentMessage":"２３ヶ月プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567024103400,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":6000,"presentMessage":"２周年記念プレゼントです！","createdAt":1661606352,"updatedAt":1661606352},{"id":1140558567024103400,"userId":1140558566977966100,"sentAt":1661606352,"itemType":1,"itemId":1,"amount":6000,"presentMessage":"２.５周年プレゼントです！","createdAt":1661606352,"updatedAt":1661606352}]}}`),
			},
			want:    []string{"sessionId", "updatedResources.user.id", "updatedResources.userCards.id", "updatedResources.userCards.userId", "updatedResources.userDecks.cardId1", "updatedResources.userDecks.cardId2", "updatedResources.userDecks.cardId3", "updatedResources.userDecks.id", "updatedResources.userDecks.userId", "updatedResources.userDevice.id", "updatedResources.userDevice.userId", "updatedResources.userLoginBonuses.id", "updatedResources.userLoginBonuses.userId", "updatedResources.userPresents.id", "updatedResources.userPresents.userId", "userId"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jsonDiffKeys(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonDiffKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonDiffKeys() got = %#v, want %v", got, tt.want)
			}
		})
	}
}
