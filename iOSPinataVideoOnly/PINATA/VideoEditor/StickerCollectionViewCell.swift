//
//  StickerCollectionViewCell.swift
//  PINATA
//
//  Created by Karim Abedrabbo on 11/18/18.
//  Copyright © 2018 Karim Abedrabbo. All rights reserved.
//
import UIKit

class StickerCollectionViewCell: UICollectionViewCell {
    var stickerImage: UIImageView = {
        let sticker = UIImageView()
        sticker.translatesAutoresizingMaskIntoConstraints = false
        return sticker
    }()

    
    override init(frame: CGRect) {
        super.init(frame: frame)
        self.addSubview(stickerImage)
        self.stickerImage.topAnchor.constraint(equalTo: self.topAnchor).isActive = true
        self.stickerImage.bottomAnchor.constraint(equalTo: self.bottomAnchor).isActive = true
        self.stickerImage.trailingAnchor.constraint(equalTo: self.trailingAnchor).isActive = true
        self.stickerImage.leadingAnchor.constraint(equalTo: self.leadingAnchor).isActive = true
    }
    
    required init?(coder aDecoder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }
    
}
