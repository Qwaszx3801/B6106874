/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
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
    EntEducationlevelEdges,
    EntEducationlevelEdgesFromJSON,
    EntEducationlevelEdgesFromJSONTyped,
    EntEducationlevelEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntEducationlevel
 */
export interface EntEducationlevel {
    /**
     * Educationlevelname holds the value of the "Educationlevelname" field.
     * @type {string}
     * @memberof EntEducationlevel
     */
    educationlevelname?: string;
    /**
     * 
     * @type {EntEducationlevelEdges}
     * @memberof EntEducationlevel
     */
    edges?: EntEducationlevelEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntEducationlevel
     */
    id?: number;
}

export function EntEducationlevelFromJSON(json: any): EntEducationlevel {
    return EntEducationlevelFromJSONTyped(json, false);
}

export function EntEducationlevelFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntEducationlevel {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'educationlevelname': !exists(json, 'Educationlevelname') ? undefined : json['Educationlevelname'],
        'edges': !exists(json, 'edges') ? undefined : EntEducationlevelEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntEducationlevelToJSON(value?: EntEducationlevel | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Educationlevelname': value.educationlevelname,
        'edges': EntEducationlevelEdgesToJSON(value.edges),
        'id': value.id,
    };
}


