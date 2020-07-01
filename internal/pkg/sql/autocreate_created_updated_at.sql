CREATE OR REPLACE FUNCTION trigger_set_updated_at()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION trigger_set_created_at()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.created_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_set_updated_at
    BEFORE UPDATE
    ON health_menstruation_daily_entries
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

CREATE TRIGGER trigger_set_created_at
    BEFORE INSERT
    ON health_menstruation_daily_entries
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_created_at();

CREATE TRIGGER trigger_set_updated_at
    BEFORE UPDATE
    ON health_menstruation_personal_infos
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

CREATE TRIGGER trigger_set_created_at
    BEFORE INSERT
    ON health_menstruation_personal_infos
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_created_at();

CREATE TRIGGER trigger_set_updated_at
    BEFORE UPDATE
    ON chat_messages
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

CREATE TRIGGER trigger_set_created_at
    BEFORE INSERT
    ON chat_messages
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_created_at();

CREATE TRIGGER trigger_set_updated_at
    BEFORE UPDATE
    ON chat_rooms
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

CREATE TRIGGER trigger_set_created_at
    BEFORE INSERT
    ON chat_rooms
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_created_at();

CREATE TRIGGER trigger_set_updated_at
    BEFORE UPDATE
    ON chat_room_participants
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

CREATE TRIGGER trigger_set_created_at
    BEFORE INSERT
    ON chat_room_participants
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_created_at();

CREATE TRIGGER trigger_set_updated_at
    BEFORE UPDATE
    ON log_activities
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

CREATE TRIGGER trigger_set_created_at
    BEFORE INSERT
    ON log_activities
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_created_at();

CREATE TRIGGER trigger_set_updated_at
    BEFORE UPDATE
    ON notification_devices
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

CREATE TRIGGER trigger_set_created_at
    BEFORE INSERT
    ON notification_devices
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_created_at();

CREATE TRIGGER trigger_set_updated_at
    BEFORE UPDATE
    ON notification_settings
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

CREATE TRIGGER trigger_set_created_at
    BEFORE INSERT
    ON notification_settings
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_created_at();

CREATE TRIGGER trigger_set_updated_at
    BEFORE UPDATE
    ON journal_entries
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

CREATE TRIGGER trigger_set_created_at
    BEFORE INSERT
    ON journal_entries
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_created_at();