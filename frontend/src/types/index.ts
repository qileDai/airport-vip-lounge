export interface APIResponse<T = any> {
  success: boolean;
  data?: T;
  message?: string;
  error?: string;
  meta?: Meta;
}

export interface Meta {
  page: number;
  page_size: number;
  total_count: number;
  total_pages: number;
}

export interface MemberBenefit {
  id: number;
  code: string;
  name: string;
  status: string;
  member_level: string;
  remaining_quota: number;
  valid_from: string;
  valid_until: string;
  owner: string;
  batch_id: string;
  airport_code: string;
  lounge_id: string;
  notes: string;
  created_at: string;
  updated_at: string;
}

export interface AppointmentRecord {
  id: number;
  code: string;
  name: string;
  status: string;
  member_benefit_id: number;
  flight_number: string;
  flight_date: string;
  scheduled_time: string;
  actual_arrival_time: string;
  companion_count: number;
  owner: string;
  batch_id: string;
  verification_result_id: number;
  notes: string;
  created_at: string;
  updated_at: string;
}

export interface FlightTimeSlot {
  id: number;
  code: string;
  name: string;
  status: string;
  airport_code: string;
  gate: string;
  slot_start: string;
  slot_end: string;
  capacity: number;
  occupied_count: number;
  owner: string;
  batch_id: string;
  notes: string;
  created_at: string;
  updated_at: string;
}

export interface Companion {
  id: number;
  code: string;
  name: string;
  status: string;
  appointment_record_id: number;
  id_type: string;
  id_number: string;
  member_level: string;
  relation_type: string;
  age: number;
  owner: string;
  batch_id: string;
  verification_status: string;
  notes: string;
  created_at: string;
  updated_at: string;
}

export interface WaitingListEntry {
  id: number;
  code: string;
  name: string;
  status: string;
  member_benefit_id: number;
  priority_score: number;
  queue_position: number;
  wait_start_time: string;
  expected_entry_time: string;
  actual_entry_time: string;
  reason: string;
  owner: string;
  batch_id: string;
  notes: string;
  created_at: string;
  updated_at: string;
}

export interface VerificationResult {
  id: number;
  code: string;
  name: string;
  status: string;
  appointment_record_id: number;
  member_benefit_id: number;
  verification_time: string;
  verifier: string;
  result: string;
  rejection_reason: string;
  rule_hits: string;
  score: number;
  owner: string;
  batch_id: string;
  notes: string;
  created_at: string;
  updated_at: string;
}

export interface ExceptionEvent {
  id: number;
  code: string;
  name: string;
  status: string;
  exception_type: string;
  severity: string;
  entity_type: string;
  entity_id: number;
  trigger_field: string;
  threshold_value: string;
  actual_value: string;
  handler: string;
  deadline: string;
  resolution: string;
  owner: string;
  batch_id: string;
  notes: string;
  created_at: string;
  updated_at: string;
}

export interface StatisticsData {
  total_verifications: number;
  passed_count: number;
  failed_count: number;
  pass_rate: number;
  waiting_list_count: number;
  admitted_from_waiting: number;
  admission_rate: number;
  active_benefits_count: number;
  usage_rate: number;
}
