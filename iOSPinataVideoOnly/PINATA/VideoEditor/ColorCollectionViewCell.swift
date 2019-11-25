//
//  ColorCollectionViewCell.swift
//  PINATA
//
//  Created by Karim Abedrabbo on 11/18/18.
//  Copyright © 2018 Karim Abedrabbo. All rights reserved.
//

import UIKit

class ColorCollectionViewCell: UICollectionViewCell {
    
    var colorView: UIView = {
        let color = UIView()
        color.translatesAutoresizingMaskIntoConstraints = false
        color.backgroundColor = .clear
        color.layer.cornerRadius = 15
        return color
    }()
    
    override init(frame: CGRect) {
        super.init(frame: frame)
        self.addSubview(colorView)
        self.colorView.centerXAnchor.constraint(equalTo: self.centerXAnchor).isActive = true
        self.colorView.centerYAnchor.constraint(equalTo: self.centerYAnchor).isActive = true
        self.colorView.heightAnchor.constraint(equalToConstant: 30).isActive = true
        self.colorView.widthAnchor.constraint(equalToConstant: 30).isActive = true
        
    }
    
    
    
    required init?(coder aDecoder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }
    
    override var isSelected: Bool {
        didSet {
            if isSelected {
                let previouTransform =  colorView.transform
                UIView.animate(withDuration: 0.2,
                               animations: {
                                self.colorView.transform = self.colorView.transform.scaledBy(x: 1.3, y: 1.3)
                },
                               completion: { _ in
                                UIView.animate(withDuration: 0.2) {
                                    self.colorView.transform  = previouTransform
                                }
                })
            } else {
                // animate deselection
            }
        }
    }
    
}
