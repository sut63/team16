/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Patient
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeFromJSONTyped,
    EntEmployeeToJSON,
    EntPayment,
    EntPaymentFromJSON,
    EntPaymentFromJSONTyped,
    EntPaymentToJSON,
    EntPromotionamount,
    EntPromotionamountFromJSON,
    EntPromotionamountFromJSONTyped,
    EntPromotionamountToJSON,
    EntPromotiontype,
    EntPromotiontypeFromJSON,
    EntPromotiontypeFromJSONTyped,
    EntPromotiontypeToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPromotionEdges
 */
export interface EntPromotionEdges {
    /**
     * 
     * @type {EntEmployee}
     * @memberof EntPromotionEdges
     */
    employee?: EntEmployee;
    /**
     * Payment holds the value of the payment edge.
     * @type {Array<EntPayment>}
     * @memberof EntPromotionEdges
     */
    payment?: Array<EntPayment>;
    /**
     * 
     * @type {EntPromotionamount}
     * @memberof EntPromotionEdges
     */
    promotionamount?: EntPromotionamount;
    /**
     * 
     * @type {EntPromotiontype}
     * @memberof EntPromotionEdges
     */
    promotiontype?: EntPromotiontype;
}

export function EntPromotionEdgesFromJSON(json: any): EntPromotionEdges {
    return EntPromotionEdgesFromJSONTyped(json, false);
}

export function EntPromotionEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPromotionEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'employee': !exists(json, 'employee') ? undefined : EntEmployeeFromJSON(json['employee']),
        'payment': !exists(json, 'payment') ? undefined : ((json['payment'] as Array<any>).map(EntPaymentFromJSON)),
        'promotionamount': !exists(json, 'promotionamount') ? undefined : EntPromotionamountFromJSON(json['promotionamount']),
        'promotiontype': !exists(json, 'promotiontype') ? undefined : EntPromotiontypeFromJSON(json['promotiontype']),
    };
}

export function EntPromotionEdgesToJSON(value?: EntPromotionEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'employee': EntEmployeeToJSON(value.employee),
        'payment': value.payment === undefined ? undefined : ((value.payment as Array<any>).map(EntPaymentToJSON)),
        'promotionamount': EntPromotionamountToJSON(value.promotionamount),
        'promotiontype': EntPromotiontypeToJSON(value.promotiontype),
    };
}


