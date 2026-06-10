import { describe, it, expect } from 'vitest';
import apiService from '../src/services/api';

describe('API Service', () => {
  describe('Member Benefits API', () => {
    it('should have getMemberBenefits method defined', () => {
      expect(apiService.getMemberBenefits).toBeDefined();
    });

    it('should have createMemberBenefit method defined', () => {
      expect(apiService.createMemberBenefit).toBeDefined();
    });

    it('should have updateMemberBenefit method defined', () => {
      expect(apiService.updateMemberBenefit).toBeDefined();
    });

    it('should have deleteMemberBenefit method defined', () => {
      expect(apiService.deleteMemberBenefit).toBeDefined();
    });

    it('should have transitionMemberBenefitStatus method defined', () => {
      expect(apiService.transitionMemberBenefitStatus).toBeDefined();
    });
  });

  describe('Appointments API', () => {
    it('should have getAppointments method defined', () => {
      expect(apiService.getAppointments).toBeDefined();
    });

    it('should have createAppointment method defined', () => {
      expect(apiService.createAppointment).toBeDefined();
    });

    it('should have transitionAppointmentStatus method defined', () => {
      expect(apiService.transitionAppointmentStatus).toBeDefined();
    });
  });

  describe('Statistics API', () => {
    it('should have getStatistics method defined', () => {
      expect(apiService.getStatistics).toBeDefined();
    });
  });

  describe('Health Check API', () => {
    it('should have healthCheck method defined', () => {
      expect(apiService.healthCheck).toBeDefined();
    });
  });

  describe('Domain Calculations API', () => {
    it('should have calculatePriority method defined', () => {
      expect(apiService.calculatePriority).toBeDefined();
    });

    it('should have checkConsistency method defined', () => {
      expect(apiService.checkConsistency).toBeDefined();
    });

    it('should have detectExpiredBenefits method defined', () => {
      expect(apiService.detectExpiredBenefits).toBeDefined();
    });
  });

  describe('Waiting List API', () => {
    it('should have getWaitingList method defined', () => {
      expect(apiService.getWaitingList).toBeDefined();
    });

    it('should have processWaitingList method defined', () => {
      expect(apiService.processWaitingList).toBeDefined();
    });
  });

  describe('Companions API', () => {
    it('should have batchCreateCompanions method defined', () => {
      expect(apiService.batchCreateCompanions).toBeDefined();
    });
  });

  describe('Seed Data API', () => {
    it('should have resetSeedData method defined', () => {
      expect(apiService.resetSeedData).toBeDefined();
    });
  });
});
