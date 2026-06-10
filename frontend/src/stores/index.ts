import { writable, derived } from 'svelte/store';
import type { MemberBenefit, AppointmentRecord, FlightTimeSlot, Companion, WaitingListEntry, VerificationResult, ExceptionEvent } from '../types';
import apiService from '../services/api';

// Member Benefits Store
function createMemberBenefitsStore() {
  const { subscribe, set, update } = writable<{
    data: MemberBenefit[];
    loading: boolean;
    error: string | null;
    selectedMemberBenefit: MemberBenefit | null;
    filters: Record<string, any>;
    pagination: { page: number; pageSize: number; totalCount: number; totalPages: number };
  }>({
    data: [],
    loading: false,
    error: null,
    selectedMemberBenefit: null,
    filters: {},
    pagination: { page: 1, pageSize: 10, totalCount: 0, totalPages: 0 }
  });

  return {
    subscribe,
    
    async fetchMemberBenefits(page: number = 1, pageSize: number = 10) {
      update(state => ({ ...state, loading: true, error: null }));
      try {
        const response = await apiService.getMemberBenefits(page, pageSize);
        if (response.success && response.data && response.meta) {
          update(state => ({
            ...state,
            data: response.data as MemberBenefit[],
            pagination: response.meta,
            loading: false
          }));
        }
      } catch (error) {
        update(state => ({
          ...state,
          error: (error as Error).message,
          loading: false
        }));
      }
    },

    async createMemberBenefit(data: Partial<MemberBenefit>) {
      update(state => ({ ...state, loading: true }));
      try {
        const response = await apiService.createMemberBenefit(data);
        if (response.success) {
          await this.fetchMemberBenefits();
          return response.data;
        }
        throw new Error(response.error || 'Failed to create member benefit');
      } catch (error) {
        update(state => ({
          ...state,
          error: (error as Error).message,
          loading: false
        }));
        throw error;
      }
    },

    async updateMemberBenefit(id: number, data: Partial<MemberBenefit>) {
      try {
        const response = await apiService.updateMemberBenefit(id, data);
        if (response.success) {
          await this.fetchMemberBenefits();
          return response.data;
        }
        throw new Error(response.error || 'Failed to update member benefit');
      } catch (error) {
        update(state => ({
          ...state,
          error: (error as Error).message
        }));
        throw error;
      }
    },

    async deleteMemberBenefit(id: number) {
      try {
        await apiService.deleteMemberBenefit(id);
        await this.fetchMemberBenefits();
      } catch (error) {
        update(state => ({
          ...state,
          error: (error as Error).message
        }));
        throw error;
      }
    },

    async transitionStatus(id: number, toStatus: string, action: string, reason?: string) {
      try {
        const response = await apiService.transitionMemberBenefitStatus(id, toStatus, action, reason);
        if (response.success) {
          await this.fetchMemberBenefits();
          return response.data;
        }
        throw new Error(response.error || 'Failed to transition status');
      } catch (error) {
        update(state => ({
          ...state,
          error: (error as Error).message
        }));
        throw error;
      }
    },

    selectMemberBenefit(benefit: MemberBenefit | null) {
      update(state => ({ ...state, selectedMemberBenefit: benefit }));
    },

    setFilters(filters: Record<string, any>) {
      update(state => ({ ...state, filters }));
    },

    reset() {
      set({
        data: [],
        loading: false,
        error: null,
        selectedMemberBenefit: null,
        filters: {},
        pagination: { page: 1, pageSize: 10, totalCount: 0, totalPages: 0 }
      });
    }
  };
}

export const memberBenefitsStore = createMemberBenefitsStore();

// Appointments Store
function createAppointmentsStore() {
  const { subscribe, set, update } = writable<{
    data: AppointmentRecord[];
    loading: boolean;
    error: string | null;
    selectedAppointment: AppointmentRecord | null;
    pagination: { page: number; pageSize: number; totalCount: number; totalPages: number };
  }>({
    data: [],
    loading: false,
    error: null,
    selectedAppointment: null,
    pagination: { page: 1, pageSize: 10, totalCount: 0, totalPages: 0 }
  });

  return {
    subscribe,

    async fetchAppointments(page: number = 1, pageSize: number = 10) {
      update(state => ({ ...state, loading: true, error: null }));
      try {
        const response = await apiService.getAppointments(page, pageSize);
        if (response.success && response.data && response.meta) {
          update(state => ({
            ...state,
            data: response.data as AppointmentRecord[],
            pagination: response.meta,
            loading: false
          }));
        }
      } catch (error) {
        update(state => ({
          ...state,
          error: (error as Error).message,
          loading: false
        }));
      }
    },

    selectAppointment(appointment: AppointmentRecord | null) {
      update(state => ({ ...state, selectedAppointment: appointment }));
    },

    reset() {
      set({
        data: [],
        loading: false,
        error: null,
        selectedAppointment: null,
        pagination: { page: 1, pageSize: 10, totalCount: 0, totalPages: 0 }
      });
    }
  };
}

export const appointmentsStore = createAppointmentsStore();

// Flight Slots Store
function createFlightSlotsStore() {
  const { subscribe, set, update } = writable<{
    data: FlightTimeSlot[];
    loading: boolean;
    error: string | null;
    selectedSlot: FlightTimeSlot | null;
  }>({
    data: [],
    loading: false,
    error: null,
    selectedSlot: null
  });

  return {
    subscribe,

    async fetchFlightSlots(page: number = 1, pageSize: number = 10) {
      update(state => ({ ...state, loading: true, error: null }));
      try {
        const response = await apiService.getFlightSlots(page, pageSize);
        if (response.success && response.data) {
          update(state => ({
            ...state,
            data: response.data as FlightTimeSlot[],
            loading: false
          }));
        }
      } catch (error) {
        update(state => ({
          ...state,
          error: (error as Error).message,
          loading: false
        }));
      }
    },

    async fetchByAirport(airportCode: string) {
      update(state => ({ ...state, loading: true }));
      try {
        const response = await apiService.getFlightSlotsByAirport(airportCode);
        if (response.success && response.data) {
          update(state => ({
            ...state,
            data: response.data as FlightTimeSlot[],
            loading: false
          }));
        }
      } catch (error) {
        update(state => ({
          ...state,
          error: (error as Error).message,
          loading: false
        }));
      }
    },

    selectSlot(slot: FlightTimeSlot | null) {
      update(state => ({ ...state, selectedSlot: slot }));
    },

    reset() {
      set({
        data: [],
        loading: false,
        error: null,
        selectedSlot: null
      });
    }
  };
}

export const flightSlotsStore = createFlightSlotsStore();

// UI Store for global state
function createUIStore() {
  const { subscribe, set, update } = writable<{
    commandPaletteOpen: boolean;
    detailDrawerOpen: boolean;
    sidebarCollapsed: boolean;
    notifications: Array<{ id: string; message: string; type: 'success' | 'error' | 'warning' | 'info' }>;
    activeTab: string;
  }>({
    commandPaletteOpen: false,
    detailDrawerOpen: false,
    sidebarCollapsed: false,
    notifications: [],
    activeTab: 'dashboard'
  });

  return {
    subscribe,

    toggleCommandPalette() {
      update(state => ({ ...state, commandPaletteOpen: !state.commandPaletteOpen }));
    },

    toggleDetailDrawer() {
      update(state => ({ ...state, detailDrawerOpen: !state.detailDrawerOpen }));
    },

    toggleSidebar() {
      update(state => ({ ...state, sidebarCollapsed: !state.sidebarCollapsed }));
    },

    addNotification(message: string, type: 'success' | 'error' | 'warning' | 'info' = 'info') {
      const id = Date.now().toString();
      update(state => ({
        ...state,
        notifications: [...state.notifications, { id, message, type }]
      }));

      setTimeout(() => {
        update(state => ({
          ...state,
          notifications: state.notifications.filter(n => n.id !== id)
        }));
      }, 5000);
    },

    setActiveTab(tab: string) {
      update(state => ({ ...state, activeTab: tab }));
    },

    reset() {
      set({
        commandPaletteOpen: false,
        detailDrawerOpen: false,
        sidebarCollapsed: false,
        notifications: [],
        activeTab: 'dashboard'
      });
    }
  };
}

export const uiStore = createUIStore();

// Statistics Store
function createStatisticsStore() {
  const { subscribe, set, update } = writable<{
    data: any;
    loading: boolean;
    error: string | null;
  }>({
    data: null,
    loading: false,
    error: null
  });

  return {
    subscribe,

    async fetchStatistics(dateFrom?: string, dateTo?: string) {
      update(state => ({ ...state, loading: true, error: null }));
      try {
        const response = await apiService.getStatistics(dateFrom, dateTo);
        if (response.success && response.data) {
          update(state => ({
            ...state,
            data: response.data,
            loading: false
          }));
        }
      } catch (error) {
        update(state => ({
          ...state,
          error: (error as Error).message,
          loading: false
        }));
      }
    },

    reset() {
      set({
        data: null,
        loading: false,
        error: null
      });
    }
  };
}

export const statisticsStore = createStatisticsStore();
