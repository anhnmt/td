// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// StatsBroadcastStats represents TL type `stats.broadcastStats#bdf78394`.
// Channel statistics¹.
//
// Links:
//  1) https://core.telegram.org/api/stats
//
// See https://core.telegram.org/constructor/stats.broadcastStats for reference.
type StatsBroadcastStats struct {
	// Period in consideration
	Period StatsDateRangeDays
	// Follower count change for period in consideration
	Followers StatsAbsValueAndPrev
	// total_viewcount/postcount, for posts posted during the period in consideration
	// (views_per_post). Note that in this case, current refers to the period in
	// consideration (min_date till max_date), and prev refers to the previous period
	// ((min_date - (max_date - min_date)) till min_date).
	ViewsPerPost StatsAbsValueAndPrev
	// total_viewcount/postcount, for posts posted during the period in consideration
	// (views_per_post). Note that in this case, current refers to the period in
	// consideration (min_date till max_date), and prev refers to the previous period
	// ((min_date - (max_date - min_date)) till min_date)
	SharesPerPost StatsAbsValueAndPrev
	// Percentage of subscribers with enabled notifications
	EnabledNotifications StatsPercentValue
	// Channel growth graph (absolute subscriber count)
	GrowthGraph StatsGraphClass
	// Followers growth graph (relative subscriber count)
	FollowersGraph StatsGraphClass
	// Muted users graph (relative)
	MuteGraph StatsGraphClass
	// Views per hour graph (absolute)
	TopHoursGraph StatsGraphClass
	// Interactions graph (absolute)
	InteractionsGraph StatsGraphClass
	// IV interactions graph (absolute)
	IvInteractionsGraph StatsGraphClass
	// Views by source graph (absolute)
	ViewsBySourceGraph StatsGraphClass
	// New followers by source graph (absolute)
	NewFollowersBySourceGraph StatsGraphClass
	// Subscriber language graph (pie chart)
	LanguagesGraph StatsGraphClass
	// Recent message interactions
	RecentMessageInteractions []MessageInteractionCounters
}

// StatsBroadcastStatsTypeID is TL type id of StatsBroadcastStats.
const StatsBroadcastStatsTypeID = 0xbdf78394

// Ensuring interfaces in compile-time for StatsBroadcastStats.
var (
	_ bin.Encoder     = &StatsBroadcastStats{}
	_ bin.Decoder     = &StatsBroadcastStats{}
	_ bin.BareEncoder = &StatsBroadcastStats{}
	_ bin.BareDecoder = &StatsBroadcastStats{}
)

func (b *StatsBroadcastStats) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Period.Zero()) {
		return false
	}
	if !(b.Followers.Zero()) {
		return false
	}
	if !(b.ViewsPerPost.Zero()) {
		return false
	}
	if !(b.SharesPerPost.Zero()) {
		return false
	}
	if !(b.EnabledNotifications.Zero()) {
		return false
	}
	if !(b.GrowthGraph == nil) {
		return false
	}
	if !(b.FollowersGraph == nil) {
		return false
	}
	if !(b.MuteGraph == nil) {
		return false
	}
	if !(b.TopHoursGraph == nil) {
		return false
	}
	if !(b.InteractionsGraph == nil) {
		return false
	}
	if !(b.IvInteractionsGraph == nil) {
		return false
	}
	if !(b.ViewsBySourceGraph == nil) {
		return false
	}
	if !(b.NewFollowersBySourceGraph == nil) {
		return false
	}
	if !(b.LanguagesGraph == nil) {
		return false
	}
	if !(b.RecentMessageInteractions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *StatsBroadcastStats) String() string {
	if b == nil {
		return "StatsBroadcastStats(nil)"
	}
	type Alias StatsBroadcastStats
	return fmt.Sprintf("StatsBroadcastStats%+v", Alias(*b))
}

// FillFrom fills StatsBroadcastStats from given interface.
func (b *StatsBroadcastStats) FillFrom(from interface {
	GetPeriod() (value StatsDateRangeDays)
	GetFollowers() (value StatsAbsValueAndPrev)
	GetViewsPerPost() (value StatsAbsValueAndPrev)
	GetSharesPerPost() (value StatsAbsValueAndPrev)
	GetEnabledNotifications() (value StatsPercentValue)
	GetGrowthGraph() (value StatsGraphClass)
	GetFollowersGraph() (value StatsGraphClass)
	GetMuteGraph() (value StatsGraphClass)
	GetTopHoursGraph() (value StatsGraphClass)
	GetInteractionsGraph() (value StatsGraphClass)
	GetIvInteractionsGraph() (value StatsGraphClass)
	GetViewsBySourceGraph() (value StatsGraphClass)
	GetNewFollowersBySourceGraph() (value StatsGraphClass)
	GetLanguagesGraph() (value StatsGraphClass)
	GetRecentMessageInteractions() (value []MessageInteractionCounters)
}) {
	b.Period = from.GetPeriod()
	b.Followers = from.GetFollowers()
	b.ViewsPerPost = from.GetViewsPerPost()
	b.SharesPerPost = from.GetSharesPerPost()
	b.EnabledNotifications = from.GetEnabledNotifications()
	b.GrowthGraph = from.GetGrowthGraph()
	b.FollowersGraph = from.GetFollowersGraph()
	b.MuteGraph = from.GetMuteGraph()
	b.TopHoursGraph = from.GetTopHoursGraph()
	b.InteractionsGraph = from.GetInteractionsGraph()
	b.IvInteractionsGraph = from.GetIvInteractionsGraph()
	b.ViewsBySourceGraph = from.GetViewsBySourceGraph()
	b.NewFollowersBySourceGraph = from.GetNewFollowersBySourceGraph()
	b.LanguagesGraph = from.GetLanguagesGraph()
	b.RecentMessageInteractions = from.GetRecentMessageInteractions()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsBroadcastStats) TypeID() uint32 {
	return StatsBroadcastStatsTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsBroadcastStats) TypeName() string {
	return "stats.broadcastStats"
}

// TypeInfo returns info about TL type.
func (b *StatsBroadcastStats) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stats.broadcastStats",
		ID:   StatsBroadcastStatsTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Period",
			SchemaName: "period",
		},
		{
			Name:       "Followers",
			SchemaName: "followers",
		},
		{
			Name:       "ViewsPerPost",
			SchemaName: "views_per_post",
		},
		{
			Name:       "SharesPerPost",
			SchemaName: "shares_per_post",
		},
		{
			Name:       "EnabledNotifications",
			SchemaName: "enabled_notifications",
		},
		{
			Name:       "GrowthGraph",
			SchemaName: "growth_graph",
		},
		{
			Name:       "FollowersGraph",
			SchemaName: "followers_graph",
		},
		{
			Name:       "MuteGraph",
			SchemaName: "mute_graph",
		},
		{
			Name:       "TopHoursGraph",
			SchemaName: "top_hours_graph",
		},
		{
			Name:       "InteractionsGraph",
			SchemaName: "interactions_graph",
		},
		{
			Name:       "IvInteractionsGraph",
			SchemaName: "iv_interactions_graph",
		},
		{
			Name:       "ViewsBySourceGraph",
			SchemaName: "views_by_source_graph",
		},
		{
			Name:       "NewFollowersBySourceGraph",
			SchemaName: "new_followers_by_source_graph",
		},
		{
			Name:       "LanguagesGraph",
			SchemaName: "languages_graph",
		},
		{
			Name:       "RecentMessageInteractions",
			SchemaName: "recent_message_interactions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *StatsBroadcastStats) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode stats.broadcastStats#bdf78394 as nil")
	}
	buf.PutID(StatsBroadcastStatsTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *StatsBroadcastStats) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode stats.broadcastStats#bdf78394 as nil")
	}
	if err := b.Period.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field period: %w", err)
	}
	if err := b.Followers.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field followers: %w", err)
	}
	if err := b.ViewsPerPost.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field views_per_post: %w", err)
	}
	if err := b.SharesPerPost.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field shares_per_post: %w", err)
	}
	if err := b.EnabledNotifications.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field enabled_notifications: %w", err)
	}
	if b.GrowthGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field growth_graph is nil")
	}
	if err := b.GrowthGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field growth_graph: %w", err)
	}
	if b.FollowersGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field followers_graph is nil")
	}
	if err := b.FollowersGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field followers_graph: %w", err)
	}
	if b.MuteGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field mute_graph is nil")
	}
	if err := b.MuteGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field mute_graph: %w", err)
	}
	if b.TopHoursGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field top_hours_graph is nil")
	}
	if err := b.TopHoursGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field top_hours_graph: %w", err)
	}
	if b.InteractionsGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field interactions_graph is nil")
	}
	if err := b.InteractionsGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field interactions_graph: %w", err)
	}
	if b.IvInteractionsGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field iv_interactions_graph is nil")
	}
	if err := b.IvInteractionsGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field iv_interactions_graph: %w", err)
	}
	if b.ViewsBySourceGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field views_by_source_graph is nil")
	}
	if err := b.ViewsBySourceGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field views_by_source_graph: %w", err)
	}
	if b.NewFollowersBySourceGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field new_followers_by_source_graph is nil")
	}
	if err := b.NewFollowersBySourceGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field new_followers_by_source_graph: %w", err)
	}
	if b.LanguagesGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field languages_graph is nil")
	}
	if err := b.LanguagesGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field languages_graph: %w", err)
	}
	buf.PutVectorHeader(len(b.RecentMessageInteractions))
	for idx, v := range b.RecentMessageInteractions {
		if err := v.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field recent_message_interactions element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *StatsBroadcastStats) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode stats.broadcastStats#bdf78394 to nil")
	}
	if err := buf.ConsumeID(StatsBroadcastStatsTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *StatsBroadcastStats) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode stats.broadcastStats#bdf78394 to nil")
	}
	{
		if err := b.Period.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field period: %w", err)
		}
	}
	{
		if err := b.Followers.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field followers: %w", err)
		}
	}
	{
		if err := b.ViewsPerPost.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field views_per_post: %w", err)
		}
	}
	{
		if err := b.SharesPerPost.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field shares_per_post: %w", err)
		}
	}
	{
		if err := b.EnabledNotifications.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field enabled_notifications: %w", err)
		}
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field growth_graph: %w", err)
		}
		b.GrowthGraph = value
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field followers_graph: %w", err)
		}
		b.FollowersGraph = value
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field mute_graph: %w", err)
		}
		b.MuteGraph = value
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field top_hours_graph: %w", err)
		}
		b.TopHoursGraph = value
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field interactions_graph: %w", err)
		}
		b.InteractionsGraph = value
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field iv_interactions_graph: %w", err)
		}
		b.IvInteractionsGraph = value
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field views_by_source_graph: %w", err)
		}
		b.ViewsBySourceGraph = value
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field new_followers_by_source_graph: %w", err)
		}
		b.NewFollowersBySourceGraph = value
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field languages_graph: %w", err)
		}
		b.LanguagesGraph = value
	}
	{
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field recent_message_interactions: %w", err)
		}

		if headerLen > 0 {
			b.RecentMessageInteractions = make([]MessageInteractionCounters, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value MessageInteractionCounters
			if err := value.Decode(buf); err != nil {
				return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field recent_message_interactions: %w", err)
			}
			b.RecentMessageInteractions = append(b.RecentMessageInteractions, value)
		}
	}
	return nil
}

// GetPeriod returns value of Period field.
func (b *StatsBroadcastStats) GetPeriod() (value StatsDateRangeDays) {
	if b == nil {
		return
	}
	return b.Period
}

// GetFollowers returns value of Followers field.
func (b *StatsBroadcastStats) GetFollowers() (value StatsAbsValueAndPrev) {
	if b == nil {
		return
	}
	return b.Followers
}

// GetViewsPerPost returns value of ViewsPerPost field.
func (b *StatsBroadcastStats) GetViewsPerPost() (value StatsAbsValueAndPrev) {
	if b == nil {
		return
	}
	return b.ViewsPerPost
}

// GetSharesPerPost returns value of SharesPerPost field.
func (b *StatsBroadcastStats) GetSharesPerPost() (value StatsAbsValueAndPrev) {
	if b == nil {
		return
	}
	return b.SharesPerPost
}

// GetEnabledNotifications returns value of EnabledNotifications field.
func (b *StatsBroadcastStats) GetEnabledNotifications() (value StatsPercentValue) {
	if b == nil {
		return
	}
	return b.EnabledNotifications
}

// GetGrowthGraph returns value of GrowthGraph field.
func (b *StatsBroadcastStats) GetGrowthGraph() (value StatsGraphClass) {
	if b == nil {
		return
	}
	return b.GrowthGraph
}

// GetFollowersGraph returns value of FollowersGraph field.
func (b *StatsBroadcastStats) GetFollowersGraph() (value StatsGraphClass) {
	if b == nil {
		return
	}
	return b.FollowersGraph
}

// GetMuteGraph returns value of MuteGraph field.
func (b *StatsBroadcastStats) GetMuteGraph() (value StatsGraphClass) {
	if b == nil {
		return
	}
	return b.MuteGraph
}

// GetTopHoursGraph returns value of TopHoursGraph field.
func (b *StatsBroadcastStats) GetTopHoursGraph() (value StatsGraphClass) {
	if b == nil {
		return
	}
	return b.TopHoursGraph
}

// GetInteractionsGraph returns value of InteractionsGraph field.
func (b *StatsBroadcastStats) GetInteractionsGraph() (value StatsGraphClass) {
	if b == nil {
		return
	}
	return b.InteractionsGraph
}

// GetIvInteractionsGraph returns value of IvInteractionsGraph field.
func (b *StatsBroadcastStats) GetIvInteractionsGraph() (value StatsGraphClass) {
	if b == nil {
		return
	}
	return b.IvInteractionsGraph
}

// GetViewsBySourceGraph returns value of ViewsBySourceGraph field.
func (b *StatsBroadcastStats) GetViewsBySourceGraph() (value StatsGraphClass) {
	if b == nil {
		return
	}
	return b.ViewsBySourceGraph
}

// GetNewFollowersBySourceGraph returns value of NewFollowersBySourceGraph field.
func (b *StatsBroadcastStats) GetNewFollowersBySourceGraph() (value StatsGraphClass) {
	if b == nil {
		return
	}
	return b.NewFollowersBySourceGraph
}

// GetLanguagesGraph returns value of LanguagesGraph field.
func (b *StatsBroadcastStats) GetLanguagesGraph() (value StatsGraphClass) {
	if b == nil {
		return
	}
	return b.LanguagesGraph
}

// GetRecentMessageInteractions returns value of RecentMessageInteractions field.
func (b *StatsBroadcastStats) GetRecentMessageInteractions() (value []MessageInteractionCounters) {
	if b == nil {
		return
	}
	return b.RecentMessageInteractions
}
