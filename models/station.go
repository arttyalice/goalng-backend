package models

import ()

// Station models
type Station struct {
  StationID                   uint32        `json:"station_id"`
  StationCode                 *string       `json:"station_code,omitempty"`
  StationName                 *string       `json:"station_name"`
  StationName2                *string       `json:"station_name2,omitempty"`
  StationName3                *string       `json:"station_name3,omitempty"`
  StationName4                *string       `json:"station_name4,omitempty"`
  StationName5                *string       `json:"station_name5,omitempty"`
  StationName6                *string       `json:"station_name6,omitempty"`
  StationName7                *string       `json:"station_name7,omitempty"`
  StationName8                *string       `json:"station_name8,omitempty"`
  QueuePrefix                 *string       `json:"queue_prefix,omitempty"`
  NoRoomQueuePrefix           *string       `json:"no_room_queue_prefix,omitempty"`
  AppointmentQueuePrefix      *string       `json:"appointment_queue_prefix,omitempty"`
  HospitalID                  uint32        `json:"hospital_id,omitempty"`
  QueueNumberType             *uint32       `json:"queue_number_type,omitempty"`
  QueueNumberIndex            *uint32       `json:"queue_number_index,omitempty"`
  QueueShowTime               *int          `json:"queue_show_time,omitempty"`
  ReasonText                  string        `json:"reason_text,omitempty"`
  StationMode                 int           `json:"station_mode,omitempty"`
  Status                      *string       `json:"status,omitempty"`
  OrderNo                     *uint32       `json:"order_no,omitempty"`
  CreatedDate                 *string       `json:"created_date,omitempty"`
  UpdatedDate                 *uint32       `json:"updated_date,omitempty"`
  StatGray                    *int          `json:"stat_gray,omitempty"`
  StatGreen                   *int          `json:"stat_green,omitempty"`
  StatYellow                  *int          `json:"stat_yellow,omitempty"`
  StatRed                     *int          `json:"stat_red,omitempty"`
  ApAllowCustQueue            *int          `json:"ap_allow_cust_queue,omitempty"`
  ApAllowCustAppoint          *int          `json:"ap_allow_cust_appoint,omitempty"`
  ApSlotQuota                 *int          `json:"ap_slot_quota,omitempty"`
  ApMinutesBeforeSubmit       *int          `json:"ap_minutes_before_submit,omitempty"`
  ApMinutesBeforeConfirm      *int          `json:"ap_minutes_before_confirm,omitempty"`
  ApRemarkMessage             *string       `json:"ap_remark_message,omitempty"`
}
