package postgres

import (
	"context"
	"pulsar/internal/core/domain/analytics"
)

func (db *Database) CreateInvocation(ctx context.Context, invocation *analytics.Invocation) error {
	result := db.conn.Create(invocation)
	return result.Error
}

func (db *Database) GetInvocationsOfLast24Hours(ctx context.Context, projectId string, status analytics.InvocationStatus) ([]*analytics.InvocationCount, error) {

	var data []*analytics.InvocationCount
	result := db.conn.Raw(`
		WITH hours AS (
			SELECT
			generate_series(
				CURRENT_TIMESTAMP - INTERVAL '23 hours',
				CURRENT_TIMESTAMP,
				INTERVAL '1 hour'
			) AS hour
		)
		SELECT
			DATE_TRUNC('hour', h.hour) as timestamp,
			COALESCE(COUNT(i.started_at), 0) AS count
		FROM
			hours h
			LEFT JOIN invocations i ON DATE_TRUNC('hour', i.started_at) = DATE_TRUNC('hour', h.hour)
			AND i.status = ? 
			AND i.project_id = ?
		GROUP BY
			h.hour
		ORDER BY
			h.hour;
	`, status, projectId).Scan(&data)

	return data, result.Error
}

func (db *Database) GetInvocationsOfLast7Days(ctx context.Context, projectId string, status analytics.InvocationStatus) ([]*analytics.InvocationCount, error) {
	var data []*analytics.InvocationCount
	result := db.conn.Raw(`
	WITH dates AS (
		SELECT
		  generate_series(
			CURRENT_DATE - INTERVAL '6 days',
			CURRENT_DATE,
			INTERVAL '1 day'
		) :: date AS day
	)
	SELECT
		d.day as timestamp,
		COALESCE(COUNT(i.started_at), 0) AS count
	FROM
		dates d
		LEFT JOIN invocations i ON DATE_TRUNC('day', i.started_at) = d.day
		AND i.status = ? 
		AND i.project_id = ?
	GROUP BY
		d.day
	ORDER BY
		d.day;
	`, status, projectId).Scan(&data)

	return data, result.Error
}
func (db *Database) GetInvocationsOfLast30Days(ctx context.Context, projectId string, status analytics.InvocationStatus) ([]*analytics.InvocationCount, error) {
	var data []*analytics.InvocationCount
	result := db.conn.Raw(`
	WITH dates AS (
		SELECT
		  generate_series(
			CURRENT_DATE - INTERVAL '29 days',
			CURRENT_DATE,
			INTERVAL '1 day'
		) :: date AS day
	)
	SELECT
		d.day as timestamp,
		COALESCE(COUNT(i.started_at), 0) AS count
	FROM
		dates d
		LEFT JOIN invocations i ON DATE_TRUNC('day', i.started_at) = d.day
		AND i.status = ? 
		AND i.project_id = ?
	GROUP BY
		d.day
	ORDER BY
		d.day;
	`, status, projectId).Scan(&data)

	return data, result.Error
}
