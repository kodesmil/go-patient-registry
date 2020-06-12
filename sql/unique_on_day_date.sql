CREATE UNIQUE INDEX year_month_day_uq ON period_daily_entries (account_id, date_trunc('day', day AT TIME ZONE 'GMT'));
