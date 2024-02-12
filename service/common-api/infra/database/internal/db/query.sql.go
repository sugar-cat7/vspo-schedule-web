// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db_sqlc

import (
	"context"
)

const getChannelsAndVideosByCreator = `-- name: GetChannelsAndVideosByCreator :many
SELECT
    cr.id, cr.name, cr.member_type,
    ch.id, ch.creator_id, ch.platform_name, ch.title, ch.description, ch.published_at, ch.total_view_count, ch.subscriber_count, ch.hidden_subscriber_count, ch.total_video_count, ch.thumbnail_url, ch.thumbnail_height, ch.thumbnail_width,
    v.id, v.channel_id, v.platform_name, v.title, v.description, v.stream_type, v.published_at, v.start_at, v.end_at, v.status, v.tags, v.view_count, v.thumbnail_url, v.thumbnail_height, v.thumbnail_width
FROM
    Creator cr
JOIN
    Channel ch ON cr.id = ch.creatorId
LEFT JOIN
    Video v ON ch.id = v.channelId
WHERE
    cr.member_type = 'vspo_jp'
LIMIT $1 OFFSET $2
`

type GetChannelsAndVideosByCreatorParams struct {
	Limit  int32
	Offset int32
}

type GetChannelsAndVideosByCreatorRow struct {
	Creator Creator
	Channel Channel
	Video   Video
}

func (q *Queries) GetChannelsAndVideosByCreator(ctx context.Context, arg GetChannelsAndVideosByCreatorParams) ([]GetChannelsAndVideosByCreatorRow, error) {
	rows, err := q.db.Query(ctx, getChannelsAndVideosByCreator, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetChannelsAndVideosByCreatorRow
	for rows.Next() {
		var i GetChannelsAndVideosByCreatorRow
		if err := rows.Scan(
			&i.Creator.ID,
			&i.Creator.Name,
			&i.Creator.MemberType,
			&i.Channel.ID,
			&i.Channel.CreatorID,
			&i.Channel.PlatformName,
			&i.Channel.Title,
			&i.Channel.Description,
			&i.Channel.PublishedAt,
			&i.Channel.TotalViewCount,
			&i.Channel.SubscriberCount,
			&i.Channel.HiddenSubscriberCount,
			&i.Channel.TotalVideoCount,
			&i.Channel.ThumbnailUrl,
			&i.Channel.ThumbnailHeight,
			&i.Channel.ThumbnailWidth,
			&i.Video.ID,
			&i.Video.ChannelID,
			&i.Video.PlatformName,
			&i.Video.Title,
			&i.Video.Description,
			&i.Video.StreamType,
			&i.Video.PublishedAt,
			&i.Video.StartAt,
			&i.Video.EndAt,
			&i.Video.Status,
			&i.Video.Tags,
			&i.Video.ViewCount,
			&i.Video.ThumbnailUrl,
			&i.Video.ThumbnailHeight,
			&i.Video.ThumbnailWidth,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
