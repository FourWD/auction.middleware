package orm

import "github.com/FourWD/auction.middleware/model"

type Vehicle struct {
	model.VehicleModel
	OpenPrice int `json:"open_price" query:"open_price" gorm:"type:int(10)"`
}

/*
CREATE VIEW `vehicles` AS
SELECT vehicles.* ,
s.name AS source_name,
b.name AS branch_name,
vt.name AS vehicle_type_name,
vm.name AS vehicle_model_name,
vb.name AS vehicle_brand_name,
vsm.name AS vehicle_submodel_name,
vdt.name AS vehicle_drive_type_name,
vg.name AS vehicle_gear_name,
vft.name AS vehicle_fuel_type_name,
vc.name AS vehicle_color_name,
vgr.name AS vehicle_grade_name,
vgr.row_order AS vehicle_grade_value,
p.name AS license_province_name,
ac.auction_id
FROM (
select * from `auction-vehicle-0001`.vehicles
UNION
select * from `auction-vehicle-0002`.vehicles
UNION
select * from `auction-vehicle-0003`.vehicles
UNION
select * from `auction-vehicle-0004`.vehicles
UNION
select * from `auction-vehicle-0005`.vehicles
UNION
select * from `auction-vehicle-0006`.vehicles
UNION
select * from `auction-vehicle-0007`.vehicles
UNION
select * from `auction-vehicle-0008`.vehicles
UNION
select * from `auction-vehicle-0009`.vehicles
UNION
select * from `auction-vehicle-0010`.vehicles
UNION
select * from `auction-vehicle-0011`.vehicles
UNION
select * from `auction-vehicle-0012`.vehicles
UNION
select * from `auction-vehicle-0013`.vehicles
UNION
select * from `auction-vehicle-0014`.vehicles
UNION
select * from `auction-vehicle-0015`.vehicles
UNION
select * from `auction-vehicle-0016`.vehicles
UNION
select * from `auction-vehicle-0017`.vehicles
UNION
select * from `auction-vehicle-0018`.vehicles
UNION
select * from `auction-vehicle-0019`.vehicles
UNION
select * from `auction-vehicle-0020`.vehicles
UNION
select * from `auction-vehicle-0021`.vehicles
UNION
select * from `auction-vehicle-0022`.vehicles
UNION
select * from `auction-vehicle-0023`.vehicles
UNION
select * from `auction-vehicle-0024`.vehicles
UNION
select * from `auction-vehicle-0025`.vehicles
UNION
select * from `auction-vehicle-0026`.vehicles
UNION
select * from `auction-vehicle-0027`.vehicles
UNION
select * from `auction-vehicle-0028`.vehicles
UNION
select * from `auction-vehicle-0029`.vehicles
UNION
select * from `auction-vehicle-0030`.vehicles
) AS vehicles
LEFT JOIN auction_vehicles ac ON ac.vehicle_id = vehicles.id
LEFT JOIN vehicle_sub_models vsm ON vsm.id = vehicles.vehicle_submodel_id
LEFT JOIN vehicle_models vm ON vm.id = vsm.vehicle_model_id
LEFT JOIN vehicle_types vt ON vt.id = vm.vehicle_type_id
LEFT JOIN vehicle_brands vb ON vb.id = vm.vehicle_brand_id
LEFT JOIN vehicle_drive_types vdt ON vdt.id = vsm.vehicle_drive_type_id
LEFT JOIN vehicle_gears vg ON vg.id = vsm.vehicle_gear_id
LEFT JOIN vehicle_fuel_types vft ON vft.id = vsm.vehicle_fuel_id
LEFT JOIN vehicle_colors vc ON vc.id = vehicles.vehicle_color_id
LEFT JOIN vehicle_grades vgr ON vgr.id = vehicles.vehicle_grade_id
LEFT JOIN sources s ON s.id = vehicles.source_id
LEFT JOIN branches b ON b.id = vehicles.branch_id
LEFT JOIN provinces p ON p.id = vehicles.license_provice_id;
*/
