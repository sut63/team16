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
    EntAge,
    EntAgeFromJSON,
    EntAgeFromJSONTyped,
    EntAgeToJSON,
    EntBookcourse,
    EntBookcourseFromJSON,
    EntBookcourseFromJSONTyped,
    EntBookcourseToJSON,
    EntEquipment,
    EntEquipmentFromJSON,
    EntEquipmentFromJSONTyped,
    EntEquipmentToJSON,
    EntEquipmentrental,
    EntEquipmentrentalFromJSON,
    EntEquipmentrentalFromJSONTyped,
    EntEquipmentrentalToJSON,
    EntPayment,
    EntPaymentFromJSON,
    EntPaymentFromJSONTyped,
    EntPaymentToJSON,
    EntPosition,
    EntPositionFromJSON,
    EntPositionFromJSONTyped,
    EntPositionToJSON,
    EntPromotion,
    EntPromotionFromJSON,
    EntPromotionFromJSONTyped,
    EntPromotionToJSON,
    EntSalary,
    EntSalaryFromJSON,
    EntSalaryFromJSONTyped,
    EntSalaryToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntEmployeeEdges
 */
export interface EntEmployeeEdges {
    /**
     * 
     * @type {EntAge}
     * @memberof EntEmployeeEdges
     */
    age?: EntAge;
    /**
     * Bookcourse holds the value of the bookcourse edge.
     * @type {Array<EntBookcourse>}
     * @memberof EntEmployeeEdges
     */
    bookcourse?: Array<EntBookcourse>;
    /**
     * Equipment holds the value of the equipment edge.
     * @type {Array<EntEquipment>}
     * @memberof EntEmployeeEdges
     */
    equipment?: Array<EntEquipment>;
    /**
     * Equipmentrental holds the value of the equipmentrental edge.
     * @type {Array<EntEquipmentrental>}
     * @memberof EntEmployeeEdges
     */
    equipmentrental?: Array<EntEquipmentrental>;
    /**
     * Payment holds the value of the payment edge.
     * @type {Array<EntPayment>}
     * @memberof EntEmployeeEdges
     */
    payment?: Array<EntPayment>;
    /**
     * 
     * @type {EntPosition}
     * @memberof EntEmployeeEdges
     */
    position?: EntPosition;
    /**
     * Promotion holds the value of the promotion edge.
     * @type {Array<EntPromotion>}
     * @memberof EntEmployeeEdges
     */
    promotion?: Array<EntPromotion>;
    /**
     * 
     * @type {EntSalary}
     * @memberof EntEmployeeEdges
     */
    salary?: EntSalary;
}

export function EntEmployeeEdgesFromJSON(json: any): EntEmployeeEdges {
    return EntEmployeeEdgesFromJSONTyped(json, false);
}

export function EntEmployeeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntEmployeeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'age': !exists(json, 'Age') ? undefined : EntAgeFromJSON(json['Age']),
        'bookcourse': !exists(json, 'Bookcourse') ? undefined : ((json['Bookcourse'] as Array<any>).map(EntBookcourseFromJSON)),
        'equipment': !exists(json, 'Equipment') ? undefined : ((json['Equipment'] as Array<any>).map(EntEquipmentFromJSON)),
        'equipmentrental': !exists(json, 'Equipmentrental') ? undefined : ((json['Equipmentrental'] as Array<any>).map(EntEquipmentrentalFromJSON)),
        'payment': !exists(json, 'Payment') ? undefined : ((json['Payment'] as Array<any>).map(EntPaymentFromJSON)),
        'position': !exists(json, 'Position') ? undefined : EntPositionFromJSON(json['Position']),
        'promotion': !exists(json, 'Promotion') ? undefined : ((json['Promotion'] as Array<any>).map(EntPromotionFromJSON)),
        'salary': !exists(json, 'Salary') ? undefined : EntSalaryFromJSON(json['Salary']),
    };
}

export function EntEmployeeEdgesToJSON(value?: EntEmployeeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'age': EntAgeToJSON(value.age),
        'bookcourse': value.bookcourse === undefined ? undefined : ((value.bookcourse as Array<any>).map(EntBookcourseToJSON)),
        'equipment': value.equipment === undefined ? undefined : ((value.equipment as Array<any>).map(EntEquipmentToJSON)),
        'equipmentrental': value.equipmentrental === undefined ? undefined : ((value.equipmentrental as Array<any>).map(EntEquipmentrentalToJSON)),
        'payment': value.payment === undefined ? undefined : ((value.payment as Array<any>).map(EntPaymentToJSON)),
        'position': EntPositionToJSON(value.position),
        'promotion': value.promotion === undefined ? undefined : ((value.promotion as Array<any>).map(EntPromotionToJSON)),
        'salary': EntSalaryToJSON(value.salary),
    };
}


