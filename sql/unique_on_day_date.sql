CREATE UNIQUE INDEX year_month_day_uq ON period_daily_entries (date_trunc('day', day AT TIME ZONE 'GMT'));
