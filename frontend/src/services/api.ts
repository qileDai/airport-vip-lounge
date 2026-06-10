import type {
  APIResponse,
  MemberBenefit,
  AppointmentRecord,
  FlightTimeSlot,
  Companion,
  WaitingListEntry,
  VerificationResult,
  ExceptionEvent,
  StatisticsData
} from '../types';

const API_BASE_URL = '/api/v1';

class APIService {
  private async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<APIResponse<T>> {
    const url = `${API_BASE_URL}${endpoint}`;
    
    const config: RequestInit = {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers,
      },
      ...options,
    };

    try {
      const response = await fetch(url, config);
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data: APIResponse<T> = await response.json();
      return data;
    } catch (error) {
      console.error('API request failed:', error);
      throw error;
    }
  }

  // Member Benefits API
  async getMemberBenefits(page: number = 1, pageSize: number = 10): Promise<APIResponse<MemberBenefit[]>> {
    return this.request<MemberBenefit[]>(`/member-benefits?page=${page}&page_size=${pageSize}`);
  }

  async getMemberBenefitById(id: number): Promise<APIResponse<MemberBenefit>> {
    return this.request<MemberBenefit>(`/member-benefits/${id}`);
  }

  async createMemberBenefit(data: Partial<MemberBenefit>): Promise<APIResponse<MemberBenefit>> {
    return this.request<MemberBenefit>('/member-benefits', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async updateMemberBenefit(id: number, data: Partial<MemberBenefit>): Promise<APIResponse<MemberBenefit>> {
    return this.request<MemberBenefit>(`/member-benefits/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  async deleteMemberBenefit(id: number): Promise<APIResponse<void>> {
    return this.request<void>(`/member-benefits/${id}`, {
      method: 'DELETE',
    });
  }

  async transitionMemberBenefitStatus(
    id: number,
    toStatus: string,
    action: string,
    reason?: string
  ): Promise<APIResponse<MemberBenefit>> {
    return this.request<MemberBenefit>(`/member-benefits/${id}/transition`, {
      method: 'POST',
      body: JSON.stringify({ to_status: toStatus, action, reason }),
    });
  }

  async getMemberBenefitsByStatus(status: string): Promise<APIResponse<MemberBenefit[]>> {
    return this.request<MemberBenefit[]>(`/member-benefits/status/${status}`);
  }

  // Appointments API
  async getAppointments(page: number = 1, pageSize: number = 10): Promise<APIResponse<AppointmentRecord[]>> {
    return this.request<AppointmentRecord[]>(`/appointments?page=${page}&page_size=${pageSize}`);
  }

  async getAppointmentById(id: number): Promise<APIResponse<AppointmentRecord>> {
    return this.request<AppointmentRecord>(`/appointments/${id}`);
  }

  async createAppointment(data: Partial<AppointmentRecord>): Promise<APIResponse<AppointmentRecord>> {
    return this.request<AppointmentRecord>('/appointments', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async updateAppointment(id: number, data: Partial<AppointmentRecord>): Promise<APIResponse<AppointmentRecord>> {
    return this.request<AppointmentRecord>(`/appointments/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  async deleteAppointment(id: number): Promise<APIResponse<void>> {
    return this.request<void>(`/appointments/${id}`, {
      method: 'DELETE',
    });
  }

  async transitionAppointmentStatus(
    id: number,
    toStatus: string,
    action: string,
    reason?: string
  ): Promise<APIResponse<AppointmentRecord>> {
    return this.request<AppointmentRecord>(`/appointments/${id}/transition`, {
      method: 'POST',
      body: JSON.stringify({ to_status: toStatus, action, reason }),
    });
  }

  // Flight Time Slots API
  async getFlightSlots(page: number = 1, pageSize: number = 10): Promise<APIResponse<FlightTimeSlot[]>> {
    return this.request<FlightTimeSlot[]>(`/flight-slots?page=${page}&page_size=${pageSize}`);
  }

  async getFlightSlotById(id: number): Promise<APIResponse<FlightTimeSlot>> {
    return this.request<FlightTimeSlot>(`/flight-slots/${id}`);
  }

  async getFlightSlotsByAirport(airportCode: string): Promise<APIResponse<FlightTimeSlot[]>> {
    return this.request<FlightTimeSlot[]>(`/flight-slots/airport/${airportCode}`);
  }

  async createFlightSlot(data: Partial<FlightTimeSlot>): Promise<APIResponse<FlightTimeSlot>> {
    return this.request<FlightTimeSlot>('/flight-slots', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  // Companions API
  async getCompanions(page: number = 1, pageSize: number = 10): Promise<APIResponse<Companion[]>> {
    return this.request<Companion[]>(`/companions?page=${page}&page_size=${pageSize}`);
  }

  async getCompanionById(id: number): Promise<APIResponse<Companion>> {
    return this.request<Companion>(`/companions/${id}`);
  }

  async getCompanionsByAppointment(appointmentId: number): Promise<APIResponse<Companion[]>> {
    return this.request<Companion[]>(`/companions/appointment/${appointmentId}`);
  }

  async createCompanion(data: Partial<Companion>): Promise<APIResponse<Companion>> {
    return this.request<Companion>('/companions', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async batchCreateCompanions(companions: Partial<Companion>[]): Promise<APIResponse<void>> {
    return this.request<void>('/companions/batch', {
      method: 'POST',
      body: JSON.stringify({ records: companions }),
    });
  }

  // Waiting List API
  async getWaitingList(page: number = 1, pageSize: number = 10): Promise<APIResponse<WaitingListEntry[]>> {
    return this.request<WaitingListEntry[]>(`/waiting-list?page=${page}&page_size=${pageSize}`);
  }

  async processWaitingList(slotId: number, availableSlots: number): Promise<APIResponse<WaitingListEntry[]>> {
    return this.request<WaitingListEntry[]>('/waiting-list/process', {
      method: 'POST',
      body: JSON.stringify({ slot_id: slotId, available_slots: availableSlots }),
    });
  }

  // Statistics API
  async getStatistics(dateFrom?: string, dateTo?: string): Promise<APIResponse<StatisticsData>> {
    const params = new URLSearchParams();
    if (dateFrom) params.append('date_from', dateFrom);
    if (dateTo) params.append('date_to', dateTo);
    
    return this.request<StatisticsData>(`/statistics?${params.toString()}`);
  }

  // Domain Calculations API
  async calculatePriority(memberBenefitId: number, flightTime?: string): Promise<APIResponse<{ priority_score: number }>> {
    const params = new URLSearchParams();
    if (flightTime) params.append('flight_time', flightTime);
    
    return this.request(`/domain-calculations/priority/${memberBenefitId}?${params.toString()}`);
  }

  async checkConsistency(appointmentId: number): Promise<APIResponse<any>> {
    return this.request(`/domain-calculations/consistency/${appointmentId}`);
  }

  async detectExpiredBenefits(): Promise<APIResponse<MemberBenefit[]>> {
    return this.request<MemberBenefit[]>('/domain-calculations/expired-benefits');
  }

  // Health Check API
  async healthCheck(): Promise<APIResponse<any>> {
    return this.request('/health');
  }

  // Seed Data API
  async resetSeedData(): Promise<APIResponse<any>> {
    return this.request('/seed/reset', {
      method: 'POST',
    });
  }

  // Exceptions API
  async getExceptions(page: number = 1, pageSize: number = 10): Promise<APIResponse<ExceptionEvent[]>> {
    return this.request<ExceptionEvent[]>(`/exceptions?page=${page}&page_size=${pageSize}`);
  }
}

export const apiService = new APIService();
export default apiService;
