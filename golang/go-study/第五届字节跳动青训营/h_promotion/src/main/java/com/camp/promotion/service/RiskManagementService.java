package com.camp.promotion.service;

import org.springframework.stereotype.Service;

@Service
public class RiskManagementService {

    public boolean riskManagement(Long userId) {
        return userId == 123456;
    }
}
