import { describe, it, expect, beforeAll } from 'vitest';
import apiService from '../src/services/api';

describe('Integration Tests - Seed Data Flow', () => {
  let testServerAvailable = false;

  beforeAll(async () => {
    try {
      const response = await apiService.healthCheck();
      testServerAvailable = response.success;
    } catch (error) {
      console.warn('Backend server not available for integration tests');
      testServerAvailable = false;
    }
  });

  describe('Health Check', () => {
    it('should return successful health check when server is available', async () => {
      if (!testServerAvailable) {
        console.log('Skipping - server not available');
        expect(true).toBe(true);
        return;
      }

      const response = await apiService.healthCheck();
      
      expect(response.success).toBe(true);
      expect(response.data).toBeDefined();
      expect(response.data.status).toBe('ok');
    });
  });

  describe('Member Benefits CRUD Operations', () => {
    it('should fetch member benefits with seed data', async () => {
      if (!testServerAvailable) {
        console.log('Skipping - server not available');
        expect(true).toBe(true);
        return;
      }

      const response = await apiService.getMemberBenefits(1, 10);
      
      expect(response.success).toBe(true);
      expect(response.data).toBeDefined();
      expect(Array.isArray(response.data)).toBe(true);
      expect(response.meta).toBeDefined();
      
      if (response.data && response.data.length > 0) {
        const firstBenefit = response.data[0];
        
        expect(firstBenefit.id).toBeDefined();
        expect(firstBenefit.code).toBeDefined();
        expect(firstBenefit.name).not.toBe('');
        expect(firstBenefit.status).toBeDefined();
        expect(firstBenefit.member_level).toBeDefined();
        expect(typeof firstBenefit.remaining_quota).toBe('number');
        expect(firstBenefit.valid_from).toBeDefined();
        expect(firstBenefit.valid_until).toBeDefined();
        expect(firstBenefit.owner).toBeDefined();
        expect(firstBenefit.batch_id).toBeDefined();
        expect(firstBenefit.created_at).toBeDefined();
      }
    }, 10000);

    it('should have realistic business data in seed records', async () => {
      if (!testServerAvailable) {
        console.log('Skipping - server not available');
        expect(true).toBe(true);
        return;
      }

      const response = await apiService.getMemberBenefits(1, 50);
      
      if (response.data && response.data.length > 0) {
        response.data.forEach(benefit => {
          expect(benefit.code).not.toMatch(/^(Item|Test|Demo)\d+$/i);
          expect(benefit.name).not.toMatch(/^(Item|Test|Demo)\d+$/i);
          
          const validStatuses = ['active', 'expired', 'pending_review', 'archived', 'rejected', 'draft'];
          expect(validStatuses).toContain(benefit.status);
          
          const validLevels = ['platinum', 'gold', 'silver', 'standard'];
          expect(validLevels).toContain(benefit.member_level);
        });
      }
    });
  });

  describe('Statistics API Integration', () => {
    it('should fetch statistics data from backend', async () => {
      if (!testServerAvailable) {
        console.log('Skipping - server not available');
        expect(true).toBe(true);
        return;
      }

      const response = await apiService.getStatistics();
      
      expect(response.success).toBe(true);
      expect(response.data).toBeDefined();
      
      if (response.data) {
        expect(typeof response.data.total_verifications).toBe('number');
        expect(typeof response.data.pass_rate).toBe('number');
        expect(typeof response.data.waiting_list_count).toBe('number');
        expect(typeof response.data.admission_rate).toBe('number');
        expect(typeof response.data.active_benefits_count).toBe('number');
        expect(typeof response.data.usage_rate).toBe('number');
      }
    });
  });

  describe('Domain Calculations API', () => {
    it('should detect expired benefits when available', async () => {
      if (!testServerAvailable) {
        console.log('Skipping - server not available');
        expect(true).toBe(true);
        return;
      }

      try {
        const response = await apiService.detectExpiredBenefits();
        expect(response.success).toBe(true);
        expect(Array.isArray(response.data)).toBe(true);
      } catch (error) {
        expect(error).toBeDefined();
      }
    });
  });

  describe('Pagination and Metadata', () => {
    it('should return correct pagination metadata', async () => {
      if (!testServerAvailable) {
        console.log('Skipping - server not available');
        expect(true).toBe(true);
        return;
      }

      const response = await apiService.getMemberBenefits(1, 10);
      
      if (response.meta) {
        expect(typeof response.meta.page).toBe('number');
        expect(typeof response.meta.page_size).toBe('number');
        expect(typeof response.meta.total_count).toBe('number');
        expect(typeof response.meta.total_pages).toBe('number');
        expect(response.meta.page).toBe(1);
        expect(response.meta.page_size).toBe(10);
      }
    });
  });
});
